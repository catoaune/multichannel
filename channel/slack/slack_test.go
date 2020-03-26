package slack_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/catoaune/multichannel/channel/slack"
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
	err := slackConfig.SendNotification("Dette er en test")
	if err != nil {
		t.Errorf("Expected nil but got %v.", err)
	}

}
