package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fbngrm/zh-mnemonics/sync"
)

// AnkiConnect API URL
const ankiConnectURL = "http://localhost:8765"

type AnkiCard struct {
	Chinese string `json:"Chinese"`
	English string `json:"English"`
	Story   string `json:"Story"`
}

var ankiCards = make(map[string]AnkiCard)

func main() {
	get("chinese::var")
}

func get(deck string) {
	filename := "anki_cards_" + deck + ".csv"
	cardIDs, err := sync.GetIDs(ankiConnectURL, deck)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through card IDs and retrieve card information
	for _, cardID := range cardIDs {
		// Get card info action
		action := "cardsInfo"
		params := map[string]interface{}{
			"cards": []interface{}{cardID},
		}

		// Perform the AnkiConnect request
		response, err := makeRequest(ankiConnectURL, action, params)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		// Parse the AnkiConnect response
		var cardInfo map[string]interface{}
		err = json.Unmarshal(response, &cardInfo)
		if err != nil {
			fmt.Println("Error parsing response:", err)
			return
		}
		results, ok := cardInfo["result"]
		if !ok {
			fmt.Println("missing result field in response")
		}
		result := results.([]interface{})[0].(map[string]interface{})

		// Extract "Story" and "Chinese" fields
		story, storyExists := result["fields"].(map[string]interface{})["Mnemonic"].(map[string]interface{})["value"].(string)
		chinese, _ := result["fields"].(map[string]interface{})["Chinese"].(map[string]interface{})["value"].(string)

		// start := `<div class="back" lang="zh-Hans">
		// <div  style="text-align:center">
		// <span class="medium japanese" style="text-align:center">`
		// end := `</span>
		// </div>
		// </div>`
		// chinese = strings.TrimLeft(chinese, start)
		// chinese = strings.TrimRight(chinese, end)
		// fmt.Println(chinese)

		// Check if "Story" field is not empty
		if storyExists && story != "" {
			// Create an AnkiCard struct and add it to the slice
			ankiCard := AnkiCard{
				Story:   story,
				Chinese: chinese,
			}
			ankiCards[chinese] = ankiCard
		}
		time.Sleep(20 * time.Millisecond)
	}

	fmt.Println(ankiCards)

	// Call the writeCSVFile function to create the CSV file
	if err := writeCSVFile(ankiCards, filename); err != nil {
		println("Error writing CSV file:", err)
		return
	}
}

func makeRequest(ankiConnectURL, action string, params map[string]interface{}) ([]byte, error) {
	// Prepare the request payload
	requestData := map[string]interface{}{
		"action":  action,
		"version": 6,
		"params":  params,
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

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

// writeCSVFile writes a CSV file from a map[string]AnkiCard.
func writeCSVFile(cards map[string]AnkiCard, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Writing header
	header := []string{"Chinese", "English", "Story"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Writing data
	for _, card := range cards {
		row := []string{card.Chinese, card.English, card.Story}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
