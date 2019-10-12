package main

import (
	"os"
	"strconv"

	"github.com/catoaune/multichannel/targetsystems/sms/gatewayapi"
)

func main() {
	key := os.Getenv("key")
	secret := os.Getenv("secret")
	msg := os.Args[1]
	receipient, _ := strconv.ParseUint(os.Args[2], 10, 64)
	gatewayapiConfig := gatewayapi.NewConfig(key, secret, "MChannel")
	gatewayapiConfig.SendNotification(msg, receipient)
}
