package pswincom

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestConfig_SendNotification(t *testing.T) {
	type fields struct {
		ConfigType string
		username   string
		password   string
		URL        string
		from       string
	}
	type args struct {
		msg       string
		recipient string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				ConfigType: tt.fields.ConfigType,
				username:   tt.fields.username,
				password:   tt.fields.password,
				URL:        tt.fields.URL,
				from:       tt.fields.from,
			}
			if err := c.SendNotification(tt.args.msg, tt.args.recipient); (err != nil) != tt.wantErr {
				t.Errorf("SendNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	type args struct {
		username string
		password string
		from     string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.username, tt.args.password, tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatNumber(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatNumber(tt.args.phoneNumber); got != tt.want {
				t.Errorf("formatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smsMessages_setTextAsHex(t *testing.T) {
	type fields struct {
		XMLName  xml.Name
		Text     string
		username string
		password string
		msglst   struct {
			Text string `xml:",chardata"`
			msg  []struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"ID"`
				message   string `xml:"TEXT"`
				sender    string `xml:"SND"`
				receiver  string `xml:"RCV"`
				operation string `xml:"OP"`
				class     string `xml:"CLASS"`
			} `xml:"MSG"`
		}
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := smsMessages{
				XMLName:  tt.fields.XMLName,
				Text:     tt.fields.Text,
				username: tt.fields.username,
				password: tt.fields.password,
				msglst:   tt.fields.msglst,
			}
		})
	}
}