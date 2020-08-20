package interpreter

import (
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

type ValExpression struct {
	val int
}

func (e *ValExpression) Interpret() int {
	return e.val
}

type PlusExpression struct {
	left  Expression
	right Expression
}

func (e *PlusExpression) Interpret() int {
	return e.left.Interpret() + e.right.Interpret()
}

type MinExpression struct {
	left  Expression
	right Expression
}

func (e *MinExpression) Interpret() int {
	return e.left.Interpret() - e.right.Interpret()
}

type Context struct {
	text string
	exp  Expression
}

func NewContext(text string) *Context {
	return &Context{text: text}
}

func (ctx *Context) Execute() int {
	ctx.parse()
	return ctx.exp.Interpret()
}

func (ctx *Context) parse() {
	chars := strings.Split(ctx.text, "")
	index := 0
	for index < len(chars) {
		v := chars[index]
		switch v {
		case "+":
			index++
			next := chars[index]
			val, _ := strconv.Atoi(next)
			ctx.exp = &PlusExpression{left: ctx.exp, right: &ValExpression{val}}
		case "-":
			index++
			next := chars[index]
			val, _ := strconv.Atoi(next)
			ctx.exp = &MinExpression{left: ctx.exp, right: &ValExpression{val}}
		default:
			val, _ := strconv.Atoi(v)
			ctx.exp = &ValExpression{val: val}
		}
		index++
	}
}
