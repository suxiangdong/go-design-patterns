package simple_factory

import (
	"errors"
)

const (
	Wechat = iota
	Ali
)

type Payment interface {
	Pay() int
}

func NewPayment(payMethod int) (Payment, error) {
	if payMethod == Wechat {
		return &wechatPay{},nil
	} else if payMethod == Ali {
		return &aliPay{},nil
	}
	return nil, errors.New("not found")
}

// 阿里支付
type aliPay struct {}

func (*aliPay) Pay() int {
	return Ali
}

// 微信支付
type wechatPay struct {}

func (*wechatPay) Pay() int {
	return Wechat
}