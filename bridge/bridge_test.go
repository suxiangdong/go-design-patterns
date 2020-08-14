package bridge

import "testing"

func TestPhone_SoftRun(t *testing.T) {
	game := &GameSoft{}
	phoneA := NewPhoneA(game)
	if r := phoneA.Run(); r != "game run\n" {
		t.Fatal("PhoneA test fail")
	}

	message := &MessageSoft{}
	phoneB := NewPhoneA(message)
	if r := phoneB.Run(); r != "message run\n" {
		t.Fatal("PhoneB test fail")
	}
}
