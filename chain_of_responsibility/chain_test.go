package chain_of_responsibility

import "testing"

func TestNewHandler_NoMatchHandler(t *testing.T) {
	manager := NewHandler(1, "主管", nil)
	if r := manager.Handle(2); r != "no handler\n" {
		t.Fatal("NoMatchHandler test fail")
	}
}

func TestNewHandler_MatchHandler(t *testing.T) {
	boss := NewHandler(2, "老板", nil)
	manager := NewHandler(1, "主管", boss)
	if r := manager.Handle(2); r != "老板\n" {
		t.Fatal("MatchHandler test fail")
	}
}
