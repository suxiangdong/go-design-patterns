package template_method

type NotificationInterface interface {
	send() string
}

type TemplateInterface interface {
	Send() string
	ParseContent() string
}

type Template struct {
	NotificationInterface
}

func (n *Template) ParseContent() string {
	return "content"
}

func (n *Template) Send() string {
	return n.send() + ":" + n.ParseContent()
}

type EmailNotification struct{ *Template }

func (n *EmailNotification) send() string {
	return "email"
}

func NewEmailNotification() *EmailNotification {
	emailNotification := &EmailNotification{}
	temp := &Template{emailNotification}
	return &EmailNotification{temp}
}

type SmsNotification struct{ *Template }

func (n *SmsNotification) send() string {
	return "sms"
}

func NewSmsNotification() *SmsNotification {
	smsNotification := &SmsNotification{}
	temp := &Template{smsNotification}
	return &SmsNotification{temp}
}
