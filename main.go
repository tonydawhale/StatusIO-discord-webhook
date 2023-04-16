package main

import (
	"os"
	"log"
	"time"
	"fmt"
	"strings"

	"StatusIO-discord-webhook/mongo"
	"StatusIO-discord-webhook/requests"
	"StatusIO-discord-webhook/types"
	"StatusIO-discord-webhook/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("Starting application...")
	loadEnv()
	mongo.Init()

	c := cron.New()
	c.AddFunc("* * * * *", check)
	c.Start()
	fmt.Scanln()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if strings.Contains(strings.ToLower(os.Getenv("DISCORD_WEBHOOK_USERNAME")), "discord") {
		log.Fatal("Webhook Username cannot contain discord")
	}
	log.Println("Succesfully Loaded .env file")
}

func updateIncident(incident types.StatusPageIncident, messageID string) {
	embed := utils.EmbedFromIncident(incident)

	if messageID == "" {
		payload := types.WebhookRequestPayload{
			Username: os.Getenv("DISCORD_WEBHOOK_USERNAME"),
			AvatarURL: os.Getenv("DISCORD_WEBHOOK_AVATAR_URL"),
			Embeds: []types.DiscordEmbed{
				embed,
			},
		}
		resp, err := requests.SendWebhookRequest(payload)
		if err != nil {
			log.Println("Error sending webhook request: " + err.Error())
			return
		}

		messageID = resp.Id
	} else {
		payload := types.WebhookRequestPayload{
			Embeds: []types.DiscordEmbed{
				embed,
			},
		}
		_, err := requests.UpdateWebhookRequest(payload, messageID)
		if err != nil {
			log.Println("Error updating webhook request: " + err.Error())
			return
		}
	}
	err := mongo.SetIncidentData(incident.Id, types.MongoIncident{
		IncidentID: incident.Id,
		LastUpdate: time.Now().Unix(),
		MessageID: messageID,
		Resolved: utils.IsResolvedStatus(incident.Status),
	})

	if err != nil {
		log.Println("Error setting incident data: " + err.Error())
		return
	}
}

func check() {
	log.Println("heartbeat")
	incidents, err := requests.FetchStatusIOData()
	if err != nil {
		log.Println(err)
		return
	}
	for _, incident := range utils.ReverseIncidents(incidents.Incidents) {
		mongoData, err := mongo.GetIncidentFromMongo(incident.Id)
		if err != nil {
			if utils.IsResolvedStatus(incident.Status) {
				continue
			}

			log.Println("New Incident: " + incident.Id)
			updateIncident(incident, "")
			return
		}

		var updateTime string
		if incident.UpdatedAt == "" {
			updateTime = incident.CreatedAt
		} else {
			updateTime = incident.UpdatedAt
		}

		incidentUpdate, err := time.Parse(time.RFC3339, updateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		if mongoData.LastUpdate < incidentUpdate.Unix() {
			log.Println("New Update: " + incident.Id)
			updateIncident(incident, mongoData.MessageID)
		}
	}
}