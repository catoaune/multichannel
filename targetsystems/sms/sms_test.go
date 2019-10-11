package sms_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/catoaune/multichannel/targetsystems/sms"
)

func TestSendNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		//quals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Write([]byte(`ok`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	//api := API{server.Client(), server.URL}
	//body, err := api.DoStuff()

	//ok(t, err)
	//equals(t, []byte("OK"), body)
	smsConfig := sms.NewConfig(server.URL)
	err := smsConfig.SendNotification("Dette er en test")
	if err != nil {
		t.Errorf("Expected nil but got %v.", err)
	}

}
