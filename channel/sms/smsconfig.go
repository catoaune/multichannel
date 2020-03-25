package sms

//ServiceConfig contains the configuration for SMS service
type ServiceConfig interface {
	SendNotification(message string, receiver string)
}
