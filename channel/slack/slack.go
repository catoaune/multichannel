// Package slack implements functions for sending messages to a Slack channel
// You need a Slack incoming webhook url to be able to messages to a Slack channel
package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

//Config for Slack
type Config struct {
	ConfigType string
	URL        string
}

//RequestBody struct for data being sent to Slack
type RequestBody struct {
	Text string `json:"text"`
}



type RequestBodyFormatted struct {
	Blocks []Blocks `json:"blocks"`
}
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Emoji bool	`json:"emoji,omitempty"`
}
type Blocks struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
	Accessory Accessory `json:"accessory,omitempty"`
}

type Accessory struct {
	Type string 	`json:"type"`
	Text Text		`json:"text"`
	Value string	`json:"value"`
}


//NewConfig returns a new Config
func NewConfig(URL string) Config {
	newConfig := Config{ConfigType: "Slack", URL: URL}
	return newConfig
}

// SendNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func (c Config) SendNotification(msg string) error {

	slackBody, _ := json.Marshal(RequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack: " + buf.String())
	}
	return nil
}

// SendFormattedNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// markdown formatted text and the slack channel is saved within Slack.
func (c Config) SendFormattedNotification(msg string) error {
	requestBodyFormatted := new(RequestBodyFormatted)
    blocks := new(Blocks)
    text := new(Text)


	text.Type = "mrkdwn"
	text.Text = msg
	blocks.Type = "section"
	blocks.Text = *text

	var block = []Blocks{}
	block = append(block, *blocks)

	requestBodyFormatted.Blocks = block
	slackBody, _ := json.Marshal(requestBodyFormatted)
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack: " + buf.String())
	}
	return nil
}

func (c Config) SendFormattedNotificationButton(msg string, button string, value string) error {
	requestBodyFormatted := new(RequestBodyFormatted)
	blocks := new(Blocks)
	text := new(Text)
	accessory := new(Accessory)


	text.Type = "mrkdwn"
	text.Text = msg
	accessory.Type = "button"
	accessory.Text.Type = "plain_text"
	accessory.Text.Text = button
	accessory.Text.Emoji = true
	accessory.Value = value
	blocks.Type = "section"
	blocks.Text = *text
	blocks.Accessory = *accessory

	var block = []Blocks{}
	block = append(block, *blocks)

	requestBodyFormatted.Blocks = block
	slackBody, _ := json.Marshal(requestBodyFormatted)
	log.Printf("JSON: %s", string(slackBody))
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack: " + buf.String())
	}
	return nil
}

// CreateMessageWithButton creates a new empty RequestBodyFormatted and return it
func (c Config) CreateMessageWithButton() *RequestBodyFormatted {
	requestBodyFormatted := new(RequestBodyFormatted)
	var block = []Blocks{}
	requestBodyFormatted.Blocks = block
	return requestBodyFormatted
}

// AddMessage adds a message with a button to the arrays of messages in RequestBodyFormatted
func (requestBodyFormatted *RequestBodyFormatted) AddMessage(msgType string, msg string, buttonLabel string, buttonValue string) {
	blocks := new(Blocks)
	text := new(Text)
	accessory := new(Accessory)


	text.Type = msgType
	text.Text = msg
	accessory.Type = "button"
	accessory.Text.Type = "plain_text"
	accessory.Text.Text = buttonLabel
	accessory.Text.Emoji = true
	accessory.Value = buttonValue
	blocks.Type = "section"
	blocks.Text = *text
	blocks.Accessory = *accessory
	requestBodyFormatted.Blocks = append(requestBodyFormatted.Blocks, *blocks)
	slackBody, _ := json.Marshal(requestBodyFormatted)
	log.Printf("AddMessage JSON: %s", string(slackBody))

}

// AddMessage adds a message with a button to the arrays of messages in RequestBodyFormatted
func (requestBodyFormatted *RequestBodyFormatted) AddMessageWithoutButton(msgType string, msg string) {
	blocks := new(Blocks)
	text := new(Text)
	text.Type = msgType
	text.Text = msg
	blocks.Type = "section"
	blocks.Text = *text
	requestBodyFormatted.Blocks = append(requestBodyFormatted.Blocks, *blocks)
	slackBody, _ := json.Marshal(requestBodyFormatted)
	log.Printf("AddMessage JSON: %s", string(slackBody))

}

// SendMessageWithButton sends the message requestBodyFormatted
func (c Config) SendMessageWithButton(requestBodyFormatted RequestBodyFormatted) error {
	slackBody, err := json.Marshal(requestBodyFormatted)
	if err != nil {
		log.Printf("Error in parsing struct to JSON: %+v", err)
	}
	log.Printf("SendMessageWithButton JSON: %s", string(slackBody))
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack: " + buf.String())
	}
	return nil
}