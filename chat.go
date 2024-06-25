package main

import (
	"context"
	"fmt"

	"log"
	"os"
	"path/filepath"
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
		log.Fatal(err)
	}
	formattedResponse := formatResponse(resp)
	log.Println(formattedResponse)

	if config.SAVE {
		// Create the "chats" directory if it doesn't exist
		chatsDir := filepath.Join(config.WD, "chats")
		if _, err := os.Stat(chatsDir); os.IsNotExist(err) {
			os.MkdirAll(chatsDir, 0755)
		}

		// Get the first 40 words of the response
		words := strings.Split(formattedResponse, " ")
		first10Words := strings.Join(words[:10], " ")

		// Construct the file name using the prompt
		fileName := fmt.Sprintf("%s.md", first10Words)

		// Write the first 40 words to the file
		filePath := filepath.Join(chatsDir, fileName)
		if err := os.WriteFile(filePath, []byte(formattedResponse), 0644); err != nil {
			// Handle the error appropriately (log, panic, etc.)
			log.Println("Error saving chat:", err)
			return
		}
		log.Println("Chat saved")

	}

	// return formattedContent
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
