package twilio

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//Config for Twilio sms service
type Config struct {
	ConfigType string
	accountSID string
	authToken  string
	URL        string
	from       string
}

//NewConfig returns a new Config
func NewConfig(accountSID string, authToken string, from string) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json", accountSID: accountSID, authToken: authToken, from: from}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(msg string, recipient string) error {
	requestData := url.Values{}
	requestData.Set("To", recipient)
	requestData.Set("From", c.from)
	requestData.Set("Body", msg)
	requestDataReader := *strings.NewReader(requestData.Encode())
	client := &http.Client{}
	request, _ := http.NewRequest("POST", c.URL, &requestDataReader)
	request.SetBasicAuth(c.accountSID, c.authToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(request)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			log.Println(data["sid"])
		}
	} else {
		return errors.New("Res: " + resp.Status)
	}
	return nil
}
