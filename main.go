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
	SAVE    bool
}

func main() {
	err := godotenv.Load()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("You do not have an .env file in this dir, make sure your linux enviroment have GEMINI_API_KEY exported, or create an env file ")
		}
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

	savegochatstr := os.Getenv("SAVE_GO_CHAT")
	savegochat := false
	if savegochatstr == "true" {
		savegochat = true
	}
	config := Config{
		API_KEY: apiKey,
		WD:      wd,
		SAVE:    savegochat,
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

	}
}

func cleanText(s string) string {
	s = strings.ToLower(s)
	words := s
	return words

}
