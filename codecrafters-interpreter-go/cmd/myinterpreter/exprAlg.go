package main

// Object algebra in Go
// https://www.tzcl.me/posts/expression-problem/

// define all operations we want to support here
type ExprAlg[Op any] interface {
	Literal(value LoxObj) Op
	Binary(lhs Op, operator TokenType, rhs Op) Op
	Unary(val Op, operator TokenType) Op
	Group(grp Op) Op
}
