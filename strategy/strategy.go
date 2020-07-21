package strategy

type PaymentStrategy interface {
	Pay() string
}

type PaymentContext struct {
	Payment PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Payment: strategy,
	}
}

func (ctx *PaymentContext) Pay() string {
	return ctx.Payment.Pay()
}

type BankPaymentStrategy struct {
	PayParams interface{} // 支付参数，与模式无关
}

func (*BankPaymentStrategy) Pay() string {
	return "bank"
}

type AliPaymentStrategy struct {
	PayParams interface{}
}

func (*AliPaymentStrategy) Pay() string {
	return "ali"
}