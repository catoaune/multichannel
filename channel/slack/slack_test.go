package slack_test

import (
	"github.com/catoaune/multichannel/channel/slack"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`ok`))
	}))
	// Close the server when test finishes
	defer server.Close()

	slackConfig := slack.NewConfig(server.URL)
	err1 := slackConfig.SendNotification("Dette er en test")
	if err1 != nil {
		t.Errorf("Expected nil but got %v.", err1)
	}
	err2 := slackConfig.SendFormattedNotification("Dette er en test")
	if err2 != nil {
		t.Errorf("Expected nil but got %v.", err2)
	}

}
