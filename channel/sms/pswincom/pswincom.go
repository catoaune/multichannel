package pswincom

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//Config for PSWinCom sms service
type Config struct {
	ConfigType string
	username string
	password  string
	URL        string
	from       string
}

type smsMessages struct {
	XMLName xml.Name `xml:"SESSION"`
	Text    string   `xml:",chardata"`
	username  string   `xml:"CLIENT"`
	password      string   `xml:"PW"`
	msglst  struct {
		Text string `xml:",chardata"`
		msg  []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"ID"`
			message string `xml:"TEXT"`
			sender  string `xml:"SND"`
			receiver  string `xml:"RCV"`
			operation   string `xml:"OP"`
			class string `xml:"CLASS"`
		} `xml:"MSG"`
	} `xml:"MSGLST"`
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
	log.Println("Req: " + requestData)
	client := &http.Client{}
	b := bytes.NewBuffer([]byte(requestData))
	request, _ := http.NewRequest("POST", c.URL, b)

	request.Header.Add("Content-Length", string(len(requestData)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	resp, _ := client.Do(request)
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New(string(resp.StatusCode) + " " + resp.Status)
	}
	return nil
}

func (sm smsMessages) setTextAsHex(text string) {

}
func formatNumber(phoneNumber string) string {
	formatted := strings.Replace(phoneNumber, "+", "", -1)
	return strings.Replace(formatted, " ", "", -1)
}

