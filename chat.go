package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func (config *Config) chat(prompt string) {
	ctx := context.Background()

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println(err)
	}
	formattedResponse := formatResponse(resp)
	log.Println(formattedResponse)
	if config.SAVE {
		saveChat(formattedResponse)
	}
	// return formattedContent
}

func saveChat(formattedResponse string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("cant save chat, error getting homedir ")
		return

	}

	currentTime := time.Now().UTC()
	// Format the current time to display just the date
	currentDate := currentTime.Format("2006-01-02")
	currentchatsdir := fmt.Sprintf("%s/gochat/data/chats/%s", homedir, currentDate)
	if _, err := os.Stat(currentchatsdir); os.IsNotExist(err) {
		if err := os.MkdirAll(currentchatsdir, 0755); err != nil {
			log.Fatalf("Error creating directory: %s", err)
		}
	}

	// Get the first 40 words of the response
	words := strings.Split(formattedResponse, " ")
	var first10Words string
	if len(words) > 10 {
		first10Words = strings.Join(words[:10], " ")
	} else {
		first10Words = strings.Join(words, " ")
	}

	// Construct the file name using the prompt
	fileName := fmt.Sprintf("%s.md", first10Words)

	// Write the words to the file

	filePath := filepath.Join(currentchatsdir, fileName)
	if err := os.WriteFile(filePath, []byte(formattedResponse), 0644); err != nil {
		log.Println("Error saving chat:", err)
		return
	}
	log.Printf("Chat saved at : %v", currentchatsdir)

}

func formatResponse(resp *genai.GenerateContentResponse) string {
	var formattedContent strings.Builder
	if resp != nil && resp.Candidates != nil {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formattedContent.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formattedContent.String()
}
