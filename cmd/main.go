package main

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse"
	"fmt"
)

func main() {
	const code = "2*8-6+5"
	lexer := lex.NewLexer(code)
	parser := parse.NewParser(lexer)
	fmt.Println(parser.Parse())
}
