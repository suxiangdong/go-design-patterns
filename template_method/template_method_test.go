package template_method

import "testing"

func TestEmailNotification_Send(t *testing.T) {
	temp := &EmailNotification{}
	if r := temp.Send(); r != "email: content" {
		t.Fatal("EmailNotification test fail")
	}
}

func TestSmsNotification_Send(t *testing.T) {
	temp := &SmsNotification{}
	if r := temp.Send(); r != "sms: content" {
		t.Fatal("SmsNotification test fail")
	}
}
