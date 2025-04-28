package main

import "fmt"

type Evaluator interface {
	Eval() LoxObj
}

// alias for f() -> f64
type EvalFunc func() LoxObj

// EvalFunc implements Evaluator
func (fn EvalFunc) Eval() LoxObj {
	return fn()
}

type EvalExpr struct{}

// ---------    implement ExprAlg[Op]

func (EvalExpr) Literal(value LoxObj) Evaluator {
	return EvalFunc(func() LoxObj {
		return value
	})
}

func (EvalExpr) Binary(lhs Evaluator, operator TokenType, rhs Evaluator) Evaluator {
	return EvalFunc(func() LoxObj {
		switch operator {
		case PLUS:
			return plus(lhs.Eval(), rhs.Eval())
		case MINUS:
			return minus(lhs.Eval(), rhs.Eval())
		case SLASH:
			return div(lhs.Eval(), rhs.Eval())
		case STAR:
			return mul(lhs.Eval(), rhs.Eval())
		case EQUAL_EQUAL:
			return equalEqual(lhs.Eval(), rhs.Eval())
		case BANG_EQUAL:
			return bangEqual(lhs.Eval(), rhs.Eval())
		case GREATER:
			return greater(lhs.Eval(), rhs.Eval())
		case GREATER_EQUAL:
			return greaterEqual(lhs.Eval(), rhs.Eval())
		case LESS:
			return less(lhs.Eval(), rhs.Eval())
		case LESS_EQUAL:
			return lessEqual(lhs.Eval(), rhs.Eval())
		default:
			panic("Unknown operation for operator")
		}
	})
}

func (EvalExpr) Group(grp Evaluator) Evaluator {
	return EvalFunc(func() LoxObj {
		return grp.Eval()
	})
}

func (EvalExpr) Unary(val Evaluator, operator TokenType) Evaluator {
	return EvalFunc(func() LoxObj {
		switch operator {
		case BANG:
			return bang(val.Eval())
		default:
			panic("Unknown Unary for operator")
		}
	})
}

// ---------  Helpers

func isSameType(objects ...LoxObj) bool {
	if len(objects) == 1 || len(objects) == 0 {
		return true
	}

	l := objects[0]
	for _, obj := range objects[1:] {
		if l.LoxType != obj.LoxType {
			return false
		}
	}
	return true
}

func isOfType(loxType LoxType, objects ...LoxObj) bool {
	for _, obj := range objects {
		if loxType != obj.LoxType {
			return false
		}
	}
	return true
}

func plus(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64)+rhs.Value.(float64), LOX_NUMBER)
	}

	if isOfType(LOX_STRING, lhs, rhs) {
		return NewLoxObj(fmt.Sprintf("%s%s", lhs.Value.(string), lhs.Value.(string)), LOX_STRING)
	}

	panic("cannot add values of different types")
}

func minus(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64)-rhs.Value.(float64), LOX_NUMBER)
	}

	panic("cannot '-' values of given types")
}

func div(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64)/rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '/' values of given types")
}

func mul(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64)*rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '*' values of given types")
}

func greater(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) > rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '>' values of given types")
}

func greaterEqual(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) >= rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '>=' values of given types")
}

func less(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) < rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '<' values of given types")
}

func lessEqual(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) <= rhs.Value.(float64), LOX_NUMBER)
	}
	panic("cannot '<=' values of given types")
}

func equalEqual(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) == rhs.Value.(float64), LOX_NUMBER)
	}

	if isOfType(LOX_STRING, lhs, rhs) {
		return NewLoxObj(lhs.Value.(string) == rhs.Value.(string), LOX_STRING)
	}

	if isOfType(LOX_BOOL, lhs, rhs) {
		return NewLoxObj(lhs.Value.(bool) == rhs.Value.(bool), LOX_BOOL)
	}

	panic("cannot '==' values of given types")
}

func bangEqual(lhs LoxObj, rhs LoxObj) LoxObj {
	if isOfType(LOX_NUMBER, lhs, rhs) {
		return NewLoxObj(lhs.Value.(float64) != rhs.Value.(float64), LOX_NUMBER)
	}

	if isOfType(LOX_STRING, lhs, rhs) {
		return NewLoxObj(lhs.Value.(string) != rhs.Value.(string), LOX_STRING)
	}

	if isOfType(LOX_BOOL, lhs, rhs) {
		return NewLoxObj(lhs.Value.(bool) != rhs.Value.(bool), LOX_BOOL)
	}

	panic("cannot '!=' values of given types")
}

func bang(value LoxObj) LoxObj {
	if value.LoxType != LOX_BOOL {
		panic("cannot negate value of given type")
	}
	return NewLoxObj(!value.Value.(bool), LOX_BOOL)
}
