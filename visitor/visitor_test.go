package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	visitors := []Visitor{&NameVisitor{}, &IdNumberVisitor{}}
	users := map[string]UserInterface{"个人123": &User{idNumber: "123", realName: "个人"}, "公司456": &Company{regNumber: "456", name: "公司"}}

	for k, user := range users {
		r := ""
		for _, visitor := range visitors {
			r += user.Accept(visitor)
		}
		if r != k {
			t.Fatal("Visitor test fail")
		}
	}
}
