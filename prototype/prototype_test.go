package prototype

import (
	"testing"
)

func TestAliPayment_Clone(t *testing.T) {
	TestData := []struct {
		account, name string
	}{
		{account: "a", name: "a"},
		{account: "b", name: "b"},
		{account: "c", name: "c"},
		{account: "d", name: "d"},
		{account: "e", name: "e"},
		{account: "f", name: "f"},
	}
	var preObj *AliPayment
	for _, v := range TestData {
		obj := new(AliPayment)
		if preObj == nil {
			obj = &AliPayment{name: v.name}
		} else {
			obj := preObj.Clone()
			obj.setAccount(v.account)
			obj.setName(v.name)
		}
		//fmt.Println(unsafe.Pointer(preObj), unsafe.Pointer(obj))
		if &preObj == &obj {
			t.Fatal("Clone test fail")
		}
		preObj = obj
	}
}
