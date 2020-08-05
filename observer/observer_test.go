package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()
	firstObserver := NewFirstObserver()
	secondObserver := NewSecondObserver()
	subject.Register(firstObserver)
	subject.Register(secondObserver)

	if r := subject.Notify(); r != "first observer\n"+
		"second observer\n" {
		t.Fatal("Observer test fail")
	}
}
