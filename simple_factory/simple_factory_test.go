package simple_factory

import "testing"

func TestWechatPay(t *testing.T) {
	payObj, err := NewPayObj(Wechat)
	if err != nil {
		t.Fatal("WechatPay not found")
	}
	if payReturn := payObj.Pay(); payReturn != Wechat {
		t.Fatal("WechatPay test fail")
	}
}

func TestAliPay(t *testing.T) {
	payObj, err := NewPayObj(Ali)
	if err != nil {
		t.Fatal("AliPay not found")
	}
	if payReturn := payObj.Pay(); payReturn != Ali {
		t.Fatal("AliPay test fail")
	}
}