package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	ctx := NewContext("1+2+3+4")
	if r := ctx.Execute(); r != 10 {
		t.Fatal("Interpreter test fail")
	}
}
