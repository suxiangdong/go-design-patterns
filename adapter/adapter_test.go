package adapter

import "testing"

func TestAliPayment_Pay(t *testing.T) {
	payment := &AliPayment{}
	if r := payment.Pay(); r != "ali pay\n" {
		t.Fatal("AliPayment test fail")
	}
}

func TestAdapter(t *testing.T) {
	payment := &BankPaymentAdapter{}
	if r := payment.Pay(); r != "bank zhi fu\n" {
		t.Fatal("Adapter test fail")
	}
}
