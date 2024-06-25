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
	API_KEY  string
	WD       string
	SAVE     bool
	SAVEPATH string
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
	fmt.Println("Welcome to GOCHAT, Your low budget copilot :) ......, You can always run help command for helps and manual ")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input your prompt ...... ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {

			continue
		}
		cleaned := cleanText(text)
		command := cleaned[0]
		switch command {

		case "help":
			// Handle the help command

			fmt.Println(`
			------Welcome to gochats , here are neccesary details.--------

			Expect more functionalities soon.

			If you set SAVE_GO_CHAT=true in env your chats will be saved as md files at $HOMEDIR/gochat/data/chats.

			tips: 
			   if you use vscode run '$HOMEDIR/gochat/data/chats' , so you can see each of your chats as an md file.

			Incoming functionality:
				add history and only create new file whe you called the tab command.
				for now this should serve as a good helper on the terminal


			`)
		default:
			// Handle other commands
			config.chat(text)
		}

	}
}

func cleanText(s string) []string {
	s = strings.ToLower(s)
	words := strings.Split(s, " ")
	return words

}
