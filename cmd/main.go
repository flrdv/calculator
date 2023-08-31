package main

import (
	"bufio"
	"calculator/backend/interpret"
	"calculator/frontend/lex"
	"calculator/frontend/parse"
	"calculator/frontend/parse/ast"
	"fmt"
	"os"
)

func input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return text
}

func repl() error {
	const prompt = "> "
	interpreter := interpret.NewInterpreter(map[string]ast.Node{
		"x": ast.Integer(5),
		"f": func(args ...ast.Node) (ast.Node, error) {
			fmt.Println(args)
			return ast.Integer(10), nil
		},
		"sum": func(args ...ast.Node) (ast.Node, error) {
			var counter ast.Integer

			for _, argNode := range args {
				arg, ok := argNode.(ast.Integer)
				if !ok {
					return nil, fmt.Errorf("cannot use %v as integer", argNode)
				}

				counter += arg
			}

			return counter, nil
		},
	})

	for {
		expr := input(prompt)
		if err := calculate(interpreter, expr); err != nil {
			fmt.Println("error:", err)
		}
	}
}

func calculate(interpreter interpret.Interpreter, expr string) error {
	tree, err := parse.NewParser(lex.NewLexer(expr)).Parse()
	if err != nil {
		return err
	}

	for _, branch := range tree {
		result, err := interpreter.Evaluate(branch)
		if err != nil {
			return err
		}

		fmt.Println(result)
	}

	return nil
}

func main() {
	//fmt.Println(calculate(interpret.NewInterpreter(nil), "-(2+2)*x"))
	fmt.Println("repl:", repl())
}
