package main

type LoxType int

const (
	LOX_OBJ LoxType = iota
	LOX_NUMBER
	LOX_STRING
	LOX_BOOL
	LOX_NIL
)

type LoxObj struct {
	LoxType LoxType
	Value   any
}

func NewLoxObj(value any, loxType LoxType) LoxObj {
	return LoxObj{Value: value, LoxType: loxType}
}
