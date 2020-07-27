package prototype

type Cloneable interface{
	Clone() Cloneable
}

type AliPayment struct {
	account, name, appId, appSecret string
}

func NewPayment(account, name string) *AliPayment {
	return &AliPayment{account: account, name: name}
}

func (p *AliPayment) Clone() *AliPayment {
	cloneP := *p
	return &cloneP
}

func (p *AliPayment) setName(name string) *AliPayment {
	p.name = name
	return p
}

func (p *AliPayment) setAccount(account string) *AliPayment {
	p.account = account
	return p
}