package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	payment := &Payment{}
	paramsDecorator := &ParamsCheckDecorator{payment}
	logDecorator := LogDecorator{paramsDecorator}
	if r := logDecorator.Pay(); r != "logcheckpay" {
		t.Fatal("Decorator test fail")
	}
}

func TestParamsCheckDecorator(t *testing.T) {
	payment := &Payment{}
	paramsDecorator := &ParamsCheckDecorator{payment}
	if r :=paramsDecorator.Pay(); r != "checkpay" {
		t.Fatal("ParamsCheckDecorator test fail")
	}
}

func TestLogDecorator(t *testing.T) {
	payment := &Payment{}
	logDecorator := &LogDecorator{payment}
	if r :=logDecorator.Pay(); r != "logpay" {
		t.Fatal("LogCheckDecorator test fail")
	}
}
