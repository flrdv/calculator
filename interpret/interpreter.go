package interpret

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse"
	"fmt"
	"math"
)

type Interpreter struct {
	names map[string]float64
}

func NewInterpreter(names map[string]float64) Interpreter {
	if names == nil {
		names = make(map[string]float64)
	}

	return Interpreter{
		names: names,
	}
}

func (i Interpreter) Execute(node parse.Node) (float64, error) {
	switch node.(type) {
	case parse.Number:
		return float64(node.(parse.Number)), nil
	case parse.ID:
		value, found := i.names[string(node.(parse.ID))]
		if !found {
			return 0, fmt.Errorf("name not found: %s", node.(parse.ID))
		}

		return value, nil
	case parse.UnOp:
		unOp := node.(parse.UnOp)
		value, err := i.Execute(unOp.Value)
		if err != nil {
			return 0, err
		}

		switch unOp.Op {
		case lex.UnPlus:
			return +value, nil
		case lex.UnMinus:
			return -value, nil
		}

		return 0, fmt.Errorf("interpreter: unknown unary: %s", unOp.Op)
	case parse.BinOp:
		binOp := node.(parse.BinOp)
		left, err := i.Execute(binOp.Left)
		if err != nil {
			return 0, err
		}

		right, err := i.Execute(binOp.Right)
		if err != nil {
			return 0, err
		}

		switch binOp.Op {
		case lex.OpPlus:
			return left + right, nil
		case lex.OpMinus:
			return left - right, nil
		case lex.OpStar:
			return left * right, nil
		case lex.OpSlash:
			return left / right, nil
		case lex.OpCaret:
			return math.Pow(left, right), nil
		}

		return 0, fmt.Errorf("interpreter: unknown operator: %s", binOp.Op)
	}

	return 0, fmt.Errorf("interpreter: unknown node: %s", node)
}
