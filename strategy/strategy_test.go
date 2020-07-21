package strategy

import "testing"

func TestBankPaymentContext(t *testing.T) {
	paymentContext := NewPaymentContext(&BankPaymentStrategy{PayParams: map[string]string{"account": "test"}})
	if payReturn := paymentContext.Pay(); payReturn != "bank" {
		t.Fatal("BankPayment test fail")
	}
}

func TestAliPaymentContext(t *testing.T) {
	paymentContext := NewPaymentContext(&AliPaymentStrategy{PayParams: map[string]string{"account": "test"}})
	if payReturn := paymentContext.Pay(); payReturn != "ali" {
		t.Fatal("AliPayment test fail")
	}
}