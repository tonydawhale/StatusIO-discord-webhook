package utils

import (
	"time"
	"strings"
	"log"
	"fmt"

	"StatusIO-discord-webhook/types"
	"StatusIO-discord-webhook/constants"
)

func IsResolvedStatus(status string) bool {
	return status == "resolved" || status == "postmortem"
}

func EmbedFromIncident(incident types.StatusPageIncident) types.DiscordEmbed {
	var color int
	if incident.Status == "resolved" || incident.Status == "postmortem" {
		color = constants.EmbedColorGreen
	} else if incident.Impact == "critical" {
		color = constants.EmbedColorRed
	} else if incident.Impact == "major" {
		color = constants.EmbedColorOrange
	} else if incident.Impact == "minor" {
		color = constants.EmbedColorYellow
	} else {
		color = constants.EmbedColorBlack
	}

	affected := []string{}

	for _, component := range incident.Components {
		affected = append(affected, component.Name)
	}

	description := append([]string{}, "• Impact: " + incident.Impact)

	if len(affected) > 0 {
		description = append(description, "• Affected: " + strings.Join(affected, ", "))
	}

	timestamp, err := time.Parse(time.RFC3339, incident.StartedAt)

	if err != nil {
		log.Println(err)
	}

	embed := types.DiscordEmbed{
		Color: color,
		Title: incident.Name,
		Timestamp: timestamp.Format(time.RFC3339),
		Footer: types.EmbedFooter{
			Text: incident.Id,	
		},
		Url: incident.Shortlink,
		Description: strings.Join(description, "\n"),
	}

	for _, update := range reverseIncidentUpdates(incident.IncidentUpdates) {
		updateDT, err := time.Parse(time.RFC3339, update.CreatedAt)

		if err != nil {
			log.Println(err)
		}
		timeString := "<t:" + fmt.Sprintf("%d", updateDT.Unix()) + ":R>"
		embed.Fields = append(embed.Fields, types.EmbedField{
			Name: "" + strings.ToUpper(update.Status[0:1]) + update.Status[1:] + " (" + timeString + ")",
			Value: update.Body,
		})
	}

	return embed
}

func reverseIncidentUpdates(arr []types.StatusPageIncidentUpdate) []types.StatusPageIncidentUpdate {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func ReverseIncidents(arr []types.StatusPageIncident) []types.StatusPageIncident {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}