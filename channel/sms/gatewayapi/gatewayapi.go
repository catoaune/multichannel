// Package gatewayapi implements functions for sending sms using GatewayAPI (www.gatewayapi.com) as service provider
// You need a GatewayAPI subscription to be able to send sms
package gatewayapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/mrjones/oauth"
)

//Config for GatewayAPI sms service
type Config struct {
	ConfigType string
	URL        string
	sender     string
	key        string
	secret     string
}

//NewConfig returns a new Config
func NewConfig(key string, secret string, sender string) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://gatewayapi.com/rest/mtsms", sender: sender, key: key, secret: secret}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(msg string, recipient uint64) error {

	consumer := oauth.NewConsumer(c.key, c.secret, oauth.ServiceProvider{})
	client, err := consumer.MakeHttpClient(&oauth.AccessToken{})
	if err != nil {
		log.Fatal(err)
	}

	// Request
	type GatewayAPIRecipient struct {
		Msisdn uint64 `json:"msisdn"`
	}
	type GatewayAPIRequest struct {
		Sender     string                `json:"sender"`
		Message    string                `json:"message"`
		Recipients []GatewayAPIRecipient `json:"recipients"`
	}
	request := &GatewayAPIRequest{
		Sender:  c.sender,
		Message: msg,
		Recipients: []GatewayAPIRecipient{
			{
				Msisdn: recipient,
			},
		},
	}

	// Send it
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		log.Fatal(err)
	}
	res, err := client.Post(
		c.URL,
		"application/json",
		&buf,
	)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		log.Fatalf("http error reply, status: %q, body: %q", res.Status, body)
		return errors.New("http error reply")
	}

	// Parse the response
	type GatewayAPIResponse struct {
		Ids []uint64
	}
	response := &GatewayAPIResponse{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		log.Fatal(err)
		return errors.New("Error mapping data")
	}
	return nil
}
