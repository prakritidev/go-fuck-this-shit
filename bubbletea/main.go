package main

import (
	"io"
	"log"
	"net/http"

	html2md "github.com/JohannesKaufmann/html-to-markdown"
	tea "github.com/charmbracelet/bubbletea"
)

// Content represents the structure for storing content
type Content struct {
	URL   string
	Title string
	Body  string
}

// Model represents the state of the Bubble Tea application
type Model struct {
	ContentMap  map[string]Content
	Query       string
	SelectedURL string
	ViewMode    string
	Markdown    string
}

// fetchHTML fetches the HTML content from the specified URL
func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

// convertHTMLToMarkdown converts HTML content to Markdown format
func convertHTMLToMarkdown(htmlContent string) (string, error) {
    converter := html2md.NewConverter("", true, nil)
    markdownContent, err := converter.ConvertString(htmlContent)
    if err != nil {
        return "", err
    }
    return markdownContent, nil
}

// Init initializes the model (no-op in this case)
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update updates the model based on messages
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "l":
			m.ViewMode = "list"
		case "d":
			m.ViewMode = "details"
		case "o":
			if _, exists := m.ContentMap[m.SelectedURL]; exists {
				m.ViewMode = "details"
				html, err := fetchHTML(m.SelectedURL)
				if err != nil {
					log.Println("Failed to fetch HTML:", err)
				} else {
					markdown, err := convertHTMLToMarkdown(html)
					if err != nil {
						log.Println("Failed to convert HTML to Markdown:", err)
					} else {
						m.Markdown = markdown
					}
				}
			}
		}
	}
	return m, nil
}

// View renders the current state of the model to a string
func (m *Model) View() string {
	var view string
	switch m.ViewMode {
	case "list":
		for url, content := range m.ContentMap {
			view += "Title: " + content.Title + "\n"
			view += "URL: " + url + "\n"
			view += "Press 'o' to open\n\n"
		}
	case "details":
		if m.Markdown != "" {
			view += m.Markdown
		}
	}
	return view
}

// fetchContent simulates fetching content and populating the ContentMap
func fetchContent(url string, contentMap map[string]Content) {
	// Simulate content fetching; this should be replaced with actual fetching logic
	contentMap[url] = Content{
		URL:   url,
		Title: "Example Title",
		Body:  "Example Body",
	}
}

func main() {
	contentMap := make(map[string]Content)
	// Fetch content and populate the map
	fetchContent("https://google.com", contentMap)

	model := &Model{ContentMap: contentMap, ViewMode: "list"}
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
