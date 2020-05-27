package linkmobility

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

//Config for GatewayAPI sms service
type Config struct {
	ConfigType string
	URL        string
	sender     string
	username   string
	password   string
	platformID string
	platformPartnerID string
}

type smsRequest struct {
	Source            string `json:"source"`
	Destination       string `json:"destination"`
	UserData          string `json:"userData"`
	PlatformID        string `json:"platformId"`
	PlatformPartnerID string `json:"platformPartnerId"`
	UseDeliveryReport bool   `json:"useDeliveryReport"`
}

type smsResponse struct {
	MessageID   string `json:"messageId"`
	ResultCode  int    `json:"resultCode"`
	Description string `json:"description"`
}

//NewConfig returns a new Config
func NewConfig(sender string, username string, password string, platformID string, platformPartnerID string ) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://gatewayapi.com/rest/mtsms", sender: sender, username: username, password: password, platformID: platformID, platformPartnerID: platformPartnerID}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(msg string, recipient string) error {
	var send smsRequest
	send.UserData = msg
	send.Destination = recipient
	send.Source = c.sender
	send.UseDeliveryReport = false
	send.PlatformID = c.platformID
	send.PlatformPartnerID = c.platformPartnerID

	req, err := json.Marshal(&send)
	if err != nil {
		log.Println("Error: ${err}")
	}
	b := bytes.NewBuffer(req)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", c.URL, b)
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(request)
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New("Res: " + resp.Status)

	}
	var data smsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err == nil {
		log.Println(string(data.ResultCode) + " " + data.Description)
	}
	return nil
}