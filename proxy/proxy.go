package proxy

type PaymentInterface interface {
	Pay() string
}

type Payment struct {}

func (p *Payment) Pay() string {
	return "pay"
}

type PaymentProxy struct {
	Payment PaymentInterface
}

func (p *PaymentProxy) check() string {
	return "check"
}

func (p *PaymentProxy) Pay() string {
	return p.check() + p.Payment.Pay()
}