package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	markdown "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	// Define your search query
	query := `"AstraZeneca Pharma" ("earnings" OR "profits" OR "losses" OR "merger" OR "acquisition" OR "buyback") -filetype:pdf -filetype:doc -advertisement -health -sports1`

	// Encode the query for URL
	searchQuery := url.QueryEscape(query)
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", searchQuery)

	// Send HTTP request to DuckDuckGo
	resp, err := http.Get(searchURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	// if resp.StatusCode != http.StatusOK {
	// 	panic(fmt.Sprintf("Failed to fetch search results: %d %s", resp.StatusCode, resp.Status))
	// }

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Convert HTML to Markdown using html-to-markdown package
	converter := markdown.NewConverter("", true, nil)
	markdownContent, err := converter.ConvertString(string(body))
	if err != nil {
		panic(err)
	}

	// Output the converted Markdown content
	// fmt.Println("Markdown Content:\n")
	fmt.Println(markdownContent)

	fmt.Println("GOOGLE Markdown Content:\n")
	google()
}
