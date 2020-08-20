package flyweight

import "testing"

func TestFactory_GetFlyweight(t *testing.T) {
	factory := NewFlyweightFactory()
	f1 := factory.GetFlyweight("a")
	f2 := factory.GetFlyweight("a")
	if f1 != f2 {
		t.Fatal("Flyweight test fail")
	}
}
