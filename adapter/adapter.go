package adapter

type IPayment interface {
	Pay() string
}

type AliPayment struct {
	IPayment
}

func (p *AliPayment) Pay() string {
	return "ali pay\n"
}

type BankPayment struct{}

func (p *BankPayment) zhiFu() string {
	return "bank zhi fu\n"
}

type BankPaymentAdapter struct {
	BankPayment
}

func (p *BankPayment) Pay() string {
	return p.zhiFu()
}
