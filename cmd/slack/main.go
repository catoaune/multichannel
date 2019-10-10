package main

import (
	"os"

	"github.com/catoaune/multichannel/targetsystems/slack"
)

func main() {
	channel := os.Getenv("slack_url")
	slack.SendNotification(channel, "Hei på deg!")
}
