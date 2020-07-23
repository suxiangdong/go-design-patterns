package decorator

// 支付接口
type PaymentInterface interface {
	Pay() string
}

// 支付类
type Payment struct {}

// 支付方法
func (p *Payment) Pay() string {
	return "pay"
}

// 如果不用组合的话，抽象出一个Decorator接口亦可。
// 参数检测装饰器
type ParamsCheckDecorator struct {
	PaymentInterface
}

func (*ParamsCheckDecorator) paramsCheck() string {
	return "check"
}

func (p *ParamsCheckDecorator) Pay() string {
	return p.paramsCheck() + p.PaymentInterface.Pay()
}

// 日志装饰器
type LogDecorator struct {
	PaymentInterface
}

func (*LogDecorator) log() string {
	return "log"
}

func (p *LogDecorator) Pay() string {
	return p.log() + p.PaymentInterface.Pay()
}