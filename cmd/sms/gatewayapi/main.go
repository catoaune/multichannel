package main

import (
	"os"

	"github.com/catoaune/multichannel/targetsystems/sms/gatewayapi"
)

func main() {
	key := os.Getenv("key")
	secret := os.Getenv("secret")
	gatewayapiConfig := gatewayapi.NewConfig(key, secret, "MultiChannel")
	gatewayapiConfig.SendNotification("Hei på deg!")
}
