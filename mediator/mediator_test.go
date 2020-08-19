package mediator

import "testing"

func TestMediator(t *testing.T) {
	lucy := &Lucy{}
	john := &John{}
	mediator := &Mediator{John: john, Lucy: lucy}
	lucy.SetMediator(mediator)
	john.SetMediator(mediator)
	if r := lucy.Talk(); r != "I am John\n" {
		t.Fatal("Lucy talk test fail")
	}
	if r := john.Talk(); r != "I am Lucy\n" {
		t.Fatal("John talk test fail")
	}
}
