package requests

import (
	"bytes"
	"errors"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"StatusIO-discord-webhook/types"
)

var httpClient = http.Client{}

func SendWebhookRequest(data types.WebhookRequestPayload) (*types.WebhookResponsePayload, error) {
	url := os.Getenv("DISCORD_WEBHOOK_URL") + "?wait=true"

	payload, err := encodePayload(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, url, payload)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("there was an error with the request to discord")
	}

	var response *types.WebhookResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

func UpdateWebhookRequest(data types.WebhookRequestPayload, messageID string) (*types.WebhookResponsePayload, error) {
	url := os.Getenv("DISCORD_WEBHOOK_URL") + "/messages/" + messageID

	payload, err := encodePayload(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPatch, url, payload)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("there was an error with the request to discord")
	}

	var response *types.WebhookResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

func FetchStatusIOData() (*types.StatusPageIncidentsResponse, error) {
	url := os.Getenv("STATUS_API_URL_BASE") + "/incidents.json"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("there was an error with the request to status.io")
	}

	var response *types.StatusPageIncidentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func encodePayload(data interface{}) (*bytes.Reader, error) {
    e, err := json.Marshal(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes.NewReader(e), nil
}