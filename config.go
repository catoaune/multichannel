package multichannel

//ConfigItem contains the configuration for one channel or service, like Slack or Twitter
type ConfigItem interface {
	SendNotification(message string)
}

//Configuration is a map containing ConfigurationItems for each available channel or service
var Configuration map[string]ConfigItem

//Config returns the current Configuration
func Config() map[string]ConfigItem {
	return Configuration
}
