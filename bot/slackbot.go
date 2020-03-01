package bot

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

/*
	CreateSlackClient sets up the slack RTM (real-time messaging) client library,
	initiating the socket connection and returning the client.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

// Got help from @tempor1s and https://dev.to/shindakun/a-simple-slack-bot-in-go---the-bot-4olg


/*
	RespondToEvents waits for messages on the Slack client's incomingEvents channel,
	and sends a response when it detects the bot has been tagged in a message with @<botTag>.
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		// Log all events
		fmt.Println("Event Received: ", msg.Type)
		
		// Switch on the incoming event type
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// The bot's prefix (@topofreddit)
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			// Get rid of the prefix
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)
			splitMessage := strings.Fields(message)

			// If they do not specify a command just @ the bot, send them the help menu.
			if message == "" {
				sendHelp(slackClient, ev.Channel)
			}

			// Basic command handler
			switch strings.ToLower(splitMessage[0]) {
			case "help":
				sendHelp(slackClient, ev.Channel)
			case "pwf":
				PikachuWhatFace(slackClient, ev.Channel)
			}

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		default:
		}
	}
}

const helpMessage = "Returns a link to the pikachu whatface meme"
const memeLink = "https://knowyourmeme.com/memes/surprised-pikachu"

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, slackChannel string) {
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// PikachuWhatFace should give you the link to the knowYourMeme page
func PikachuWhatFace(slackClient *slack.RTM, slackChannel string) {
	slackClient.SendMessage(slackClient.NewOutgoingMessage(memeLink, slackChannel))
}