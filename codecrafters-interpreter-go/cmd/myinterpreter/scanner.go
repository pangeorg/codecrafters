package main

import (
	"fmt"
	"os"
	"strings"
)

type UnterminatedStringError struct {
	line int
}

func (e UnterminatedStringError) Error() string {
	return fmt.Sprintf("[line %d] Error: Unterminated string.", e.line)
}

type UnexpectedTokenError struct {
	line int
	c    string
}

func (e UnexpectedTokenError) Error() string {
	c := e.c
	if c == "%" {
		c = "%%"
	}
	return fmt.Sprintf("[line %d] Error: Unexpected character: ", e.line) + c
}

func isDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNum(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

type Scanner struct {
	content string
	start   int
	current int
	line    int
	pos     int
}

func NewScanner(content string) *Scanner {
	return &Scanner{content: content, start: 0, current: 0, line: 1, pos: 1}
}

func isAtEnd(scanner *Scanner) bool {
	return scanner.current >= len(scanner.content)
}

func peek(scanner *Scanner) byte {
	if isAtEnd(scanner) {
		return 0
	}
	return scanner.content[scanner.current]
}

func peekNext(scanner *Scanner) byte {
	if scanner.current+1 >= len(scanner.content) {
		return 0
	}
	return scanner.content[scanner.current+1]
}

func advance(scanner *Scanner) byte {
	c := scanner.content[scanner.current]
	scanner.current++
	scanner.pos++
	return c
}

func match(scanner *Scanner, expected byte) bool {
	if isAtEnd(scanner) {
		return false
	}

	if scanner.content[scanner.current] != expected {
		return false
	}

	scanner.current++
	return true
}

func increaseLine(scanner *Scanner) {
	scanner.line = scanner.line + 1
	scanner.pos = 1
}

func scanString(scanner *Scanner) (*Token, error) {
	for peek(scanner) != '"' && !isAtEnd(scanner) {
		if peek(scanner) == '\n' {
			increaseLine(scanner)
		}
		advance(scanner)
	}

	if isAtEnd(scanner) {
		return nil, UnterminatedStringError{scanner.line}
	}
	// closing "
	advance(scanner)
	value := scanner.content[scanner.start+1 : scanner.current-1]

	return NewToken(STRING, "\""+value+"\"", value, scanner.line, scanner.pos), nil
}

func scanNumber(scanner *Scanner) (*Token, error) {
	for isDigit(peek(scanner)) {
		advance(scanner)
	}
	if peek(scanner) == '.' && isDigit(peekNext(scanner)) {
		// Consume the "."
		advance(scanner)

		for isDigit(peek(scanner)) {
			advance(scanner)
		}
	}

	text := scanner.content[scanner.start:scanner.current]

	return NewToken(NUMBER, text, formatNumber(text), scanner.line, scanner.pos), nil
}

func formatNumber(str string) string {
	if !strings.Contains(str, ".") {
		str = str + ".0"
	} else {
		trailing_zeros := 0
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] == '0' && str[i-1] != '.' {
				trailing_zeros++
			}
			if str[i] == '.' {
				break
			}
		}
		if trailing_zeros > 0 {
			str = str[:(int(len(str)) - trailing_zeros)]
		}
	}
	return str
}

func scanIdentifier(scanner *Scanner) (*Token, error) {
	for isAlphaNum(peek(scanner)) {
		advance(scanner)
	}

	text := scanner.content[scanner.start:scanner.current]

	if tokenType, ok := KEYWORDS[text]; !ok {
		return NewToken(IDENTIFIER, text, "null", scanner.line, scanner.pos), nil
	} else {
		return NewToken(tokenType, text, "null", scanner.line, scanner.pos), nil
	}
}

func createToken(scanner *Scanner, tokenType TokenType) *Token {
	return NewToken(
		tokenType,
		strings.Replace(scanner.content[scanner.start:scanner.current], "\n", "", 1),
		"null", scanner.line, scanner.pos)
}

func parseToken(scanner *Scanner) (*Token, error) {
	c := advance(scanner)
	switch c {
	case '(':
		return createToken(scanner, LEFT_PAREN), nil
	case ')':
		return createToken(scanner, RIGHT_PAREN), nil
	case '[':
		return createToken(scanner, LEFT_BRACKET), nil
	case ']':
		return createToken(scanner, RIGHT_BRACKET), nil
	case '{':
		return createToken(scanner, LEFT_BRACE), nil
	case '}':
		return createToken(scanner, RIGHT_BRACE), nil
	case ',':
		return createToken(scanner, COMMA), nil
	case '.':
		return createToken(scanner, DOT), nil
	case '-':
		return createToken(scanner, MINUS), nil
	case '+':
		return createToken(scanner, PLUS), nil
	case ';':
		return createToken(scanner, SEMICOLON), nil
	case '*':
		return createToken(scanner, STAR), nil
	case '!':
		if match(scanner, '=') {
			return createToken(scanner, BANG_EQUAL), nil
		} else {
			return createToken(scanner, BANG), nil
		}
	case '=':
		if match(scanner, '=') {
			return createToken(scanner, EQUAL_EQUAL), nil
		} else {
			return createToken(scanner, EQUAL), nil
		}
	case '<':
		if match(scanner, '=') {
			return createToken(scanner, LESS_EQUAL), nil
		} else {
			return createToken(scanner, LESS), nil
		}
	case '>':
		if match(scanner, '=') {
			return createToken(scanner, GREATER_EQUAL), nil
		} else {
			return createToken(scanner, GREATER), nil
		}
	case '/':
		if match(scanner, '/') {
			for peek(scanner) != '\n' && !isAtEnd(scanner) {
				advance(scanner)
			}
			return nil, nil
		} else {
			return createToken(scanner, SLASH), nil
		}
	case ' ', '\t':
		return nil, nil
	case '\n':
		increaseLine(scanner)
		return nil, nil
	case '"':
		return scanString(scanner)
	default:
		if isDigit(c) {
			return scanNumber(scanner)
		} else if isAlpha(c) {
			return scanIdentifier(scanner)
		} else {
			return nil, UnexpectedTokenError{scanner.line, string(c)}
		}
	}
}

func (scanner *Scanner) Run(verbose bool) ([]Token, int) {

	exit_code := 0
	var tokens []Token

	for !isAtEnd(scanner) {
		scanner.start = scanner.current
		t, err := parseToken(scanner)

		if err != nil {
			exit_code = 65
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}

		if t == nil {
			continue
		}

		if verbose {
			t.Print()
		}

		tokens = append(tokens, *t)
	}

	t := NewToken(EOF, "", "", scanner.line, scanner.pos)
	if verbose {
		t.Print()
	}
	tokens = append(tokens, *t)

	return tokens, exit_code
}
