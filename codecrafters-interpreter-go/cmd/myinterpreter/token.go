package main

import (
	"fmt"
	"os"
)

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	STAR
	DOT
	COMMA
	PLUS
	MINUS
	SEMICOLON
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	SLASH
	TAB
	SPACE
	NEWLINE
	DOUBLE_QUOTES
	AND
	CLASS
	ELSE
	FALSE
	FOR
	FUN
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	STRING
	WHILE
	IDENTIFIER
	NUMBER
	EOF
)

var KEYWORDS = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"string": STRING,
	"while":  WHILE,
}

func (t TokenType) ToSymbol() string {
	switch t {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case LEFT_BRACKET:
		return "["
	case RIGHT_BRACKET:
		return "]"
	case STAR:
		return "*"
	case DOT:
		return "."
	case COMMA:
		return ","
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case SEMICOLON:
		return ";"
	case EQUAL:
		return "="
	case BANG_EQUAL:
		return "!="
	case EQUAL_EQUAL:
		return "=="
	case BANG:
		return "!"
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case SLASH:
		return "/"
	case TAB:
		return "\\t"
	case SPACE:
		return "\\s"
	case NEWLINE:
		return "\\n"
	case DOUBLE_QUOTES:
		return "\""
	case AND:
		return "and"
	case CLASS:
		return "class"
	case ELSE:
		return "else"
	case FALSE:
		return "false"
	case FOR:
		return "for"
	case FUN:
		return "fun"
	case IF:
		return "if"
	case NIL:
		return "nil"
	case OR:
		return "or"
	case PRINT:
		return "print"
	case RETURN:
		return "return"
	case SUPER:
		return "super"
	case THIS:
		return "this"
	case TRUE:
		return "true"
	case VAR:
		return "var"
	case STRING:
		return "string"
	case WHILE:
		return "while"
	case IDENTIFIER:
		return "identifier"
	case NUMBER:
		return "number"
	case EOF:
		return "eof"
	default:
		panic(fmt.Sprintf("TokenType.ToString not able to evaluate for %v", TokenType(int(t))))
	}
}

func (t TokenType) ToString() string {
	switch t {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case LEFT_BRACKET:
		return "LEFT_BRACKET"
	case RIGHT_BRACKET:
		return "RIGHT_BRACKET"
	case STAR:
		return "STAR"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SEMICOLON:
		return "SEMICOLON"
	case EQUAL:
		return "EQUAL"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case BANG:
		return "BANG"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case SLASH:
		return "SLASH"
	case TAB:
		return "TAB"
	case SPACE:
		return "SPACE"
	case NEWLINE:
		return "NEWLINE"
	case DOUBLE_QUOTES:
		return "DOUBLE_QUOTES"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FOR:
		return "FOR"
	case FUN:
		return "FUN"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case STRING:
		return "STRING"
	case WHILE:
		return "WHILE"
	case IDENTIFIER:
		return "IDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case EOF:
		return "EOF"
	default:
		panic(fmt.Sprintf("TokenType.ToString not able to evaluate for %v", TokenType(int(t))))
	}
}

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   string
	Line      int
	Pos       int
}

func NewToken(t TokenType, lexeme string, literal string, line int, pos int) *Token {
	return &Token{TokenType: t, Lexeme: lexeme, Literal: literal, Line: line, Pos: pos}
}

func (t *Token) ToString() string {
	s := "null"
	if t.Literal != "" {
		s = t.Literal
	}
	return t.TokenType.ToString() + " " + t.Lexeme + " " + s
}

func (t *Token) Print() {
	fmt.Fprintf(os.Stdout, "%s\n", t.ToString())
}
