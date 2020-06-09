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

type Msg  struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
	Message string `xml:"TEXT"`
	Sender  string `xml:"SND"`
	Receiver  string `xml:"RCV"`
	Operation   string `xml:"OP"`
	Class string `xml:"CLASS"`
}

type Msglst struct {
	Text string `xml:",chardata"`
	Msg `xml:"MSG"`
}

type smsMessages struct {
	XMLName xml.Name `xml:"SESSION"`
	Text    string   `xml:",chardata"`
	Username  string   `xml:"CLIENT"`
	Password      string   `xml:"PW"`
	Msglst `xml:"MSGLST"`
}

//NewConfig returns a new Config
func NewConfig(username string, password string, from string) Config {
	newConfig := Config{ConfigType: "SMS", URL: "https://xml.pswin.com", username: username, password: password, from: from}
	return newConfig
}

//SendNotification sends msg to recipient as SMS
func (c Config) SendNotification(message string, recipient string) error {

	mess := Msg{
		Message:   textAsHex(message),
		Sender:    c.from,
		Receiver:  formatNumber(recipient),
		Operation: "9",
	}
	list := Msglst{
		Msg:  mess,
	}
	sms := smsMessages{Username: c.username, Password: c.password, Msglst: list}

//	requestData := "USER=" + c.username
//	requestData += "&PW=" + c.password
//	requestData += "&RCV=" + formatNumber(recipient)
//	requestData += "&SND=" + c.from
//	requestData += "&TXT=" + url.QueryEscape(msg)

	client := &http.Client{}
//	b := bytes.NewBuffer([]byte(requestData))
	x, _ := xml.Marshal(sms)
	log.Printf("Req: %v", string(x))
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

