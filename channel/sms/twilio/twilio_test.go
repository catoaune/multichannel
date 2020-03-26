package twilio_test

import (
	"github.com/catoaune/multichannel/channel/sms/twilio"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		//quals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.WriteHeader(200)
		rw.Write([]byte(`{"id":"SM12345"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	//api := API{server.Client(), server.URL}
	//body, err := api.DoStuff()

	//ok(t, err)
	//equals(t, []byte("OK"), body)
	twilioConfig := twilio.NewConfig("key", "secret", "SMS sender")
	twilioConfig.URL = server.URL // Override with test URL
	err := twilioConfig.SendNotification("This is a test", "1234567890")
	if err != nil {
		t.Errorf("Expected nil but got %v.", err)
	}

}
