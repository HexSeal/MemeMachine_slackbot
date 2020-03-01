// DO NOT EDIT THIS FILE. This is a fully complete implementation.
package main

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/HexSeal/mememachine_slackbot/bot"
)

func main() {
	port := ":" + os.Getenv("PORT")
	go http.ListenAndServe(port, nil)
	slackIt()
}

// slackIt is a function that initializes the Slackbot.
func slackIt() {
	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := bot.CreateSlackClient(botToken)
	bot.RespondToEvents(slackClient)
}