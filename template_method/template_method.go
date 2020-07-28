package template_method

type NotificationInterface interface {
	Send() string
	ParseContent() string
}

type BaseNotification struct {
	NotificationInterface
}

func (n *BaseNotification) ParseContent() string {
	return "content"
}

func (n *BaseNotification) Send() string {
	return ""
}

type EmailNotification struct{ *BaseNotification }

func (n *EmailNotification) Send() string {
	return "email: " + n.ParseContent()
}

type SmsNotification struct{ *BaseNotification }

func (n *SmsNotification) Send() string {
	return "sms: " + n.ParseContent()
}
