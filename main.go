package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY string
	WD      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Println("empty API_KEY")
		return

	}
	config := Config{
		API_KEY: apiKey,
		WD:      wd,
	}
	fmt.Println("Welcome to GOCHAT, Your low budget copilot :) ...... ")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input your prompt ...... ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {

			continue
		}
		cleaned := cleanText(text)
		prompt := cleaned

		config.chat(prompt)
		log.Println("Chat generate")

	}
}

func cleanText(s string) string {
	s = strings.ToLower(s)
	words := s
	return words

}