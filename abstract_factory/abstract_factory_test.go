package abstract_factory

import (
	"testing"
)

func TestMacFactory_CreatePhone(t *testing.T) {
	factory := NewMacFactory()
	phone := factory.CreatePhone()
	if r := phone.Call(); r != "Mac phone call\n" {
		t.Fatal("Mac phone test Fail")
	}
}

func TestMacComputer_Work(t *testing.T) {
	factory := NewMacFactory()
	computer := factory.CreateComputer()
	if r := computer.Work(); r != "Mac computer work\n" {
		t.Fatal("Mac computer test Fail")
	}
}

func TestHuaWeiFactory_CreatePhone(t *testing.T) {
	factory := NewHuaWeiFactory()
	phone := factory.CreatePhone()
	if r := phone.Call(); r != "HuaWei phone call\n" {
		t.Fatal("HuaWei phone test Fail")
	}
}

func TestHuaWeiComputer_Work(t *testing.T) {
	factory := NewHuaWeiFactory()
	computer := factory.CreateComputer()
	if r := computer.Work(); r != "HuaWei computer work\n" {
		t.Fatal("HuaWei computer test Fail")
	}
}

func TestFactorySwitch(t *testing.T) {
	// switch to mac
	factories := []IFactory{NewMacFactory(), NewHuaWeiFactory()}
	for _, factory := range factories {
		phone := factory.CreatePhone()
		phone.Call()
		computer := factory.CreateComputer()
		computer.Work()
	}
}
