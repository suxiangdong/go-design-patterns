package singleton

import "testing"

func TestSingleton(t *testing.T) {
	t.Parallel()
	if single := GetInstance(); single == nil {
		t.Fatal("Singleton test fail")
	}
}

func TestAsyncSingleton(t *testing.T) {
	t.Parallel()
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 != s2 {
		t.Fatal("Async Singleton test fail")
	}
}
