package proxy

import "testing"

func TestPaymentProxy(t *testing.T) {
	payment := &PaymentProxy{&Payment{}}
	if r := payment.Pay(); r != "checkpay" {
		t.Fatal("Proxy test fail")
	}
}
