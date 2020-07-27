package factory_method

import "testing"

func TestAliPayment(t *testing.T) {
	aliFactory := &AliFactory{}
	aliPayment := aliFactory.NewPayment()
	if r := aliPayment.Pay(); r != "ali" {
		t.Fatal("AliPayment test fail")
	}
}

func TestBankPayment(t *testing.T) {
	bankFactory := &BankFactory{}
	bankPayment := bankFactory.NewPayment()
	if r:= bankPayment.Pay(); r!= "bank" {
		t.Fatal("BankPayment test fail")
	}
}

func TestMultiCall(t *testing.T) {
	// 如果改为银行支付，仅需修改此处即可
	// 如果是简单工厂模式，调用方式为：
	// p1 = Factory{}.NewPayment('ali');
	// p2 = Factory{}.NewPayment('ali');
	// p3 = Factory{}.NewPayment('ali');
	// 这部分为重复代码，这样写是不好的
	aliFactory := AliFactory{}
	payment_1 := aliFactory.NewPayment()
	payment_2 := aliFactory.NewPayment()
	payment_3 := aliFactory.NewPayment()
	payment_1.Pay()
	payment_2.Pay()
	payment_3.Pay()
}