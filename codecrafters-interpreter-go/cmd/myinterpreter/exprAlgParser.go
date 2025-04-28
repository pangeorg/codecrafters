package main

import (
	"errors"
	"fmt"
	"strconv"
)

// parse the RNF
// expression     → equality ;
// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
// term           → factor ( ( "-" | "+" ) factor )* ;
// factor         → unary ( ( "/" | "*" ) unary )* ;
// unary          → ( "!" | "-" ) unary
//                | primary ;
// primary        → NUMBER | STRING | "true" | "false" | "nil"
//                | "(" expression ")" ;

// | Grammar notation | Code representation
// | --------------------------------------
// | Terminal         | Code to match and consume a token
// | Nonterminal      | Call to that rule’s function
// | |                | if or switch statement
// | * or +           | while or for loop
// | ?                | if statement

type ExprAlgParser[A any] struct {
	tokens  []Token
	current int
	alg     ExprAlg[A]
}

func NewParser[A any](tokens []Token, alg ExprAlg[A]) ExprAlgParser[A] {
	return ExprAlgParser[A]{tokens: tokens, current: 0, alg: alg}
}

func (p *ExprAlgParser[A]) isAtEnd() bool {
	return p.current >= len(p.tokens)-1 || p.tokens[p.current].TokenType == EOF
}

func (p *ExprAlgParser[A]) peekNext() (Token, error) {
	if p.current+1 >= len(p.tokens) {
		return Token{}, errors.New("No more tokens")
	}
	return p.tokens[p.current+1], nil
}

func (p *ExprAlgParser[A]) peek() Token {
	return p.tokens[p.current]
}

func (p *ExprAlgParser[A]) advance() Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p *ExprAlgParser[A]) previous() Token {
	return p.tokens[p.current-1]
}

func (p *ExprAlgParser[A]) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *ExprAlgParser[A]) match(tokenTypes ...TokenType) bool {
	for _, tpe := range tokenTypes {
		if p.check(tpe) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *ExprAlgParser[A]) consume(tpe TokenType, message string) (Token, error) {
	if p.check(tpe) {
		return p.advance(), nil
	}
	return Token{}, errors.New(message)
}

func (p *ExprAlgParser[A]) Parse() (A, error) {
	return p.expression()
}

// expression     → equality ;
func (p *ExprAlgParser[A]) expression() (A, error) {
	return p.equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (p *ExprAlgParser[A]) equality() (A, error) {
	expr, err := p.comparison()
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right, _ := p.comparison()
		expr = p.alg.Binary(expr, operator.TokenType, right)
	}
	return expr, err
}

// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (p *ExprAlgParser[A]) comparison() (A, error) {
	expr, err := p.term()
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right, _ := p.term()
		expr = p.alg.Binary(expr, operator.TokenType, right)
	}
	return expr, err
}

// term           → factor ( ( "-" | "+" ) factor )* ;
func (p *ExprAlgParser[A]) term() (A, error) {
	expr, err := p.factor()
	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right, _ := p.factor()
		expr = p.alg.Binary(expr, operator.TokenType, right)
	}
	return expr, err
}

// factor         → unary ( ( "/" | "*" ) unary )* ;
func (p *ExprAlgParser[A]) factor() (A, error) {
	expr, err := p.unary()
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right, _ := p.unary()
		expr = p.alg.Binary(expr, operator.TokenType, right)
	}
	return expr, err
}

// unary          → ( "!" | "-" ) unary | primary ;
func (p *ExprAlgParser[A]) unary() (A, error) {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right, err := p.unary()
		return p.alg.Unary(right, operator.TokenType), err
	}
	return p.primary()
}

// primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
func (p *ExprAlgParser[A]) primary() (A, error) {
	if p.match(FALSE) {
		return p.alg.Literal(NewLoxObj(false, LOX_BOOL)), nil
	}
	if p.match(TRUE) {
		return p.alg.Literal(NewLoxObj(true, LOX_BOOL)), nil
	}
	if p.match(NIL) {
		return p.alg.Literal(NewLoxObj(nil, LOX_NIL)), nil
	}
	if p.match(NUMBER) {
		v, _ := strconv.ParseFloat(p.previous().Literal, 64)
		return p.alg.Literal(NewLoxObj(v, LOX_NUMBER)), nil
	}
	if p.match(STRING) {
		return p.alg.Literal(NewLoxObj(p.previous().Literal, LOX_STRING)), nil
	}
	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		_, err = p.consume(RIGHT_PAREN, fmt.Sprintf("[line %d] Error at '%s': Unmatched parentheses.", p.peek().Line, p.peek().Literal))
		if err != nil {
			return p.alg.Literal(NewLoxObj("", LOX_STRING)), err
		}
		return p.alg.Group(expr), err
	}
	return p.alg.Literal(NewLoxObj("", LOX_STRING)), errors.New(fmt.Sprintf("[line %d] Error at '%s', Expect expression", p.peek().Line, p.peek().Literal))
}
