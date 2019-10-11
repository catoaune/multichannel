package sms

//Config for SMS
type Config struct {
	ConfigType string
	URL        string
}

//RequestBody struct for data being sent to Slack
type RequestBody struct {
	Text string `json:"text"`
}

//NewConfig returns a new Config
func NewConfig(URL string) Config {
	newConfig := Config{ConfigType: "SMS", URL: URL}
	return newConfig
}

func (c Config) SendNotification(msg string) error {
	return nil
}
