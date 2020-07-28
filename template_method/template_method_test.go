package template_method

import "testing"

func TestEmailNotification_Send(t *testing.T) {
	temp := NewEmailNotification()
	if r := temp.Send(); r != "email:content" {
		t.Fatal("EmailNotification test fail")
	}
}

func TestSmsNotification_Send(t *testing.T) {
	temp := NewSmsNotification()
	if r := temp.Send(); r != "sms:content" {
		t.Fatal("SmsNotification test fail")
	}
}
