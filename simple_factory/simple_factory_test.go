package simple_factory

import "testing"

func TestWechatPay(t *testing.T) {
	payment, err := NewPayment(Wechat)
	if err != nil {
		t.Fatal("WechatPay not found")
	}
	if payReturn := payment.Pay(); payReturn != Wechat {
		t.Fatal("WechatPay test fail")
	}
}

func TestAliPay(t *testing.T) {
	payment, err := NewPayment(Ali)
	if err != nil {
		t.Fatal("AliPay not found")
	}
	if payReturn := payment.Pay(); payReturn != Ali {
		t.Fatal("AliPay test fail")
	}
}