package facade

import "testing"

func TestWithoutFacade(t *testing.T) {
	task := &Task{}
	if r := task.Save(); r != "task saved\n" {
		t.Fatal("Task save (without facade) test fail")
	}

	power := &Power{}
	if r := power.Close(); r != "power closed\n" {
		t.Fatal("Power close (without facade) test fail")
	}
}

func TestSystem_TurnDown(t *testing.T) {
	system := NewFacade()
	if r := system.TurnDown(); r != "task saved\n"+
		"power closed\n" {
		t.Fatal("Facade turn down test fail")
	}
}
