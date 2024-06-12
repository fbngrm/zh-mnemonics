package sync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIDs(ankiConnectURL, deck string) ([]any, error) {
	// Prepare the request payload
	requestData := map[string]any{
		"action":  "findCards",
		"version": 6,
		"params": map[string]any{
			"query": fmt.Sprintf("deck:%s", deck),
		},
	}

	// Convert the payload to JSON
	payload, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	// Send the HTTP POST request to AnkiConnect
	resp, err := http.Post(ankiConnectURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	var responseBody map[string]any
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	// Check if the request was successful
	if err, ok := responseBody["error"]; ok && err != nil {
		fmt.Println("AnkiConnect error:", err)
		return nil, err.(error)
	}

	// Extract card IDs from the response
	cardIDs, ok := responseBody["result"].([]any)
	if !ok {
		fmt.Println("Error extracting card IDs from response")
		return nil, err
	}

	return cardIDs, nil
}
