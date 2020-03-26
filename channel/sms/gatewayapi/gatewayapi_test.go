package gatewayapi_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/catoaune/multichannel/channel/sms/gatewayapi"
)

func TestSendNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		//quals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		type GatewayAPIResponse struct {
			Ids []uint64 `json:"messageID"`
		}
		response := &GatewayAPIResponse{
			[]uint64{0},
		}
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(response); err != nil {
			log.Fatal(err)
		}
		rw.Write(buf.Bytes())
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	//api := API{server.Client(), server.URL}
	//body, err := api.DoStuff()

	//ok(t, err)
	//equals(t, []byte("OK"), body)
	gatewayapiConfig := gatewayapi.NewConfig("key", "secret", "SMS sender")
	gatewayapiConfig.URL = server.URL // Override with test URL
	err := gatewayapiConfig.SendNotification("This is a test", 1234567890)
	if err != nil {
		t.Errorf("Expected nil but got %v.", err)
	}

}
