package main

import (
	"bufio"
	"calculator/frontend/lex"
	"calculator/frontend/parse"
	"calculator/interpret"
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
	interpreter := interpret.NewInterpreter(map[string]float64{
		"x": 5,
	})

	for {
		expr := input(prompt)
		if err := calculate(interpreter, expr); err != nil {
			fmt.Println("error:", err)
		}
	}
}

func calculate(interpreter interpret.Interpreter, expr string) error {
	ast, err := parse.NewParser(lex.NewLexer(expr)).Parse()
	if err != nil {
		return err
	}

	for _, branch := range ast {
		result, err := interpreter.Execute(branch)
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
