package pswincom

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

//Config for PSWinCom sms service
type Config struct {
	ConfigType string
	username string
	password  string
	URL        string
	from       string
}


//NewConfig returns a new Config
func NewConfig(username string, password string, from string) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://simple.pswin.com", username: username, password: password, from: from}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(msg string, recipient string) error {
	requestData := "USER=" + c.username
	requestData += "&PW=" + c.password
	requestData += "&RCV=" + formatNumber(recipient)
	requestData += "&SND=" + c.from
	requestData += "&TXT=" + url.QueryEscape(msg)
	rData := StringToAsciiBytes(requestData)
	log.Println("Req: " + requestData)
	client := &http.Client{}
	b := bytes.NewBuffer([]byte(requestData))
	request, _ := http.NewRequest("POST", c.URL, b)

	request.Header.Add("Content-Length", string(len(rData)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	resp, _ := client.Do(request)
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New(string(resp.StatusCode) + " " + resp.Status)
	}
	return nil
}

func formatNumber(phoneNumber string) string {
	formatted := strings.Replace(phoneNumber, "+", "", -1)
	return strings.Replace(formatted, " ", "", -1)
}


func StringToAsciiBytes(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		i++
	}
	return t
}