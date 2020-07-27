package factory_method

// 工厂接口
type PaymentFactoryInterface interface {
	NewPayment() PaymentInterface
}

// 支付类接口
type PaymentInterface interface{
	Pay() string
}

// 支付宝工厂
type AliFactory struct {}

func (p *AliFactory) NewPayment() PaymentInterface {
	return &aliPayment{}
}

type aliPayment struct {}

func (p *aliPayment) Pay() string {
	return "ali"
}

// 银行支付工厂
type BankFactory struct {}

func (p *BankFactory) NewPayment() PaymentInterface {
	return &bankPayment{}
}

type bankPayment struct {}

func (p *bankPayment) Pay() string {
	return "bank"
}
