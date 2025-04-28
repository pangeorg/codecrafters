package main

import (
	"fmt"
	"os"
)

const (
	TOKENIZE = "tokenize"
	PARSE    = "parse"
)

func runTokenizer() int {
	filename := os.Args[2]
	raw, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	content := string(raw)
	scanner := NewScanner(content)
	_, exitCode := scanner.Run(true)
	return exitCode
}

func runParser() int {
	filename := os.Args[2]
	raw, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	content := string(raw)
	scanner := NewScanner(content)
	tokens, exitCode := scanner.Run(false)

	parser := NewParser(tokens, PrintExpr{})
	expr, err := parser.Parse()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(expr.Eval())

	return exitCode
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	exitCode := 0

	command := os.Args[1]

	switch command {
	case TOKENIZE:
		exitCode = runTokenizer()
	case PARSE:
		exitCode = runParser()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		exitCode = 1
	}

	os.Exit(exitCode)
}
