package memento

import "testing"

func TestMemento(t *testing.T) {
	gameRole := &GameRole{hp: 100, mp: 100}
	c := &Caretaker{}
	c.Set(gameRole.CreateMemento())

	gameRole.HitBoss()

	if gameRole.hp != 50 || gameRole.mp != 25 {
		t.Fatal("Memento test fail")
	}

	gameRole.ReinstateMemento(c.Get())

	if gameRole.hp != 100 || gameRole.mp != 100 {
		t.Fatal("Memento test fail")
	}
}
