package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	receiver := NewTv()
	turnOn := NewTurnOnCommand(receiver)
	turnDown := NewTurnDownCommand(receiver)

	invoker := NewInvoker()

	invoker.SetCommand(turnDown)
	if r := invoker.ExecuteCommand(); r != "turn down\n" {
		t.Fatal("TurnDownCommand test fail")
	}

	invoker.SetCommand(turnOn)
	if r := invoker.ExecuteCommand(); r != "turn on\n" {
		t.Fatal("TurnOnCommand test fail")
	}
}
