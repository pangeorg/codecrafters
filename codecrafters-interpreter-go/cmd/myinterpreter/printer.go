package main

import (
	"fmt"
	"strconv"
)

type Printer interface {
	Eval() string
}

type PrintFunc func() string

func (fn PrintFunc) Eval() string {
	return fn()
}

type PrintExpr struct{}

func (PrintExpr) Literal(value LoxObj) Printer {
	return PrintFunc(func() string {
		switch value.LoxType {
		case LOX_NUMBER:
			v := strconv.FormatFloat(value.Value.(float64), 'f', -1, 64)
			return formatNumber(v)
		case LOX_STRING:
			return fmt.Sprintf("%s", value.Value.(string))
		case LOX_BOOL:
			return strconv.FormatBool(value.Value.(bool))
		case LOX_NIL:
			return "nil"
		default:
			panic("Unknown operation for operator")
		}
	})
}

func (PrintExpr) Binary(lhs Printer, operator TokenType, rhs Printer) Printer {
	return PrintFunc(func() string {
		return fmt.Sprintf("(%s %s %s)", operator.ToSymbol(), lhs.Eval(), rhs.Eval())
	})
}

func (PrintExpr) Unary(value Printer, operator TokenType) Printer {
	return PrintFunc(func() string {
		return fmt.Sprintf("(%s %s)", operator.ToSymbol(), value.Eval())
	})
}

func (PrintExpr) Group(grp Printer) Printer {
	return PrintFunc(func() string {
		return fmt.Sprintf("(group %s)", grp.Eval())
	})
}
