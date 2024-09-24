package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiKey = "AIzaSyC-_SsjhHu_HzN84XPe9iGoRUfvsjncm7E"
const cx = "270eae0ecd082483f"

func google() {
	// Define the search query
	query := `"AstraZeneca Pharma" ("earnings" OR "profits" OR "losses" OR "merger" OR "acquisition" OR "buyback") -filetype:pdf -filetype:doc -advertisement -health -sports1`

	// Build the API request URL
	searchURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&key=%s&cx=%s", url.QueryEscape(query), apiKey, cx)

	// Send HTTP request to Google Custom Search API
	resp, err := http.Get(searchURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	// Create a slice to store all text results
	var textArray []string

	// Output the search results
	items := result["items"].([]interface{})
	for _, item := range items {
		searchResult := item.(map[string]interface{})
		snippet := searchResult["snippet"].(string)

		// Append each snippet to the text array
		textArray = append(textArray, snippet)
	}

	// Create the payload for the POST request
	payload := map[string]interface{}{
		"text": textArray,
	}

	// Marshal the payload to JSON format
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Print the JSON payload (for debugging or sending in POST request)
	fmt.Println("JSON Payload for POST request:")
	fmt.Println(string(payloadJSON))

	// You can send this JSON in a POST request to your API here
	// Example: sendPostRequest(payloadJSON)
}