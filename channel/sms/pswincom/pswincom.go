package pswincom

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
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

type msg  struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
	message string `xml:"TEXT"`
	sender  string `xml:"SND"`
	receiver  string `xml:"RCV"`
	operation   string `xml:"OP"`
	class string `xml:"CLASS"`
}

type msglst struct {
	Text string `xml:",chardata"`
	msg `xml:"MSG"`
}

type smsMessages struct {
	XMLName xml.Name `xml:"SESSION"`
	Text    string   `xml:",chardata"`
	username  string   `xml:"CLIENT"`
	password      string   `xml:"PW"`
	msglst `xml:"MSGLST"`
}

//NewConfig returns a new Config
func NewConfig(username string, password string, from string) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://simple.pswin.com", username: username, password: password, from: from}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(message string, recipient string) error {

	mess := msg{
		message:   textAsHex(message),
		sender:    c.from,
		receiver:  formatNumber(recipient),
		operation: "9",
		class:     "3",
	}
	list := msglst{
		msg:  mess,
	}
	sms := smsMessages{username: c.username, password: c.password, msglst: list}

//	requestData := "USER=" + c.username
//	requestData += "&PW=" + c.password
//	requestData += "&RCV=" + formatNumber(recipient)
//	requestData += "&SND=" + c.from
//	requestData += "&TXT=" + url.QueryEscape(msg)
	log.Printf("Req: %v", sms)
	client := &http.Client{}
//	b := bytes.NewBuffer([]byte(requestData))
	x, _ := xml.Marshal(sms)
	b := bytes.NewBuffer([]byte(x))
	request, _ := http.NewRequest("POST", c.URL, b)

	request.Header.Add("Content-Length", string(len(x)))
	request.Header.Add("Content-Type", "text/xml")

	resp, _ := client.Do(request)
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New(string(resp.StatusCode) + " " + resp.Status)
	}
	return nil
}

func textAsHex(text string) string {
	var hexStr string
	for _, runeValue := range text {
		hexStr += fmt.Sprintf("%U", runeValue)
	}
	return strings.ReplaceAll(hexStr, "U+", "")
}

func formatNumber(phoneNumber string) string {
	formatted := strings.Replace(phoneNumber, "+", "", -1)
	return strings.Replace(formatted, " ", "", -1)
}

