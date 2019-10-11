package main

import (
	"os"

	"github.com/catoaune/multichannel/targetsystems/slack"
)

func main() {
	channel := os.Getenv("slack_url")
	slackConfig := slack.NewConfig(channel)
	slackConfig.SendNotification("Hei på deg!")
}
