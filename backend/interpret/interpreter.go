package interpret

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse/ast"
	"fmt"
	"math"
	"reflect"
)

type Interpreter struct {
	names map[string]ast.Node
}

func NewInterpreter(names map[string]ast.Node) Interpreter {
	if names == nil {
		names = make(map[string]ast.Node)
	}

	return Interpreter{
		names: names,
	}
}

func (i Interpreter) Evaluate(node ast.Node) (ast.Node, error) {
	switch node.(type) {
	case ast.Integer, ast.Float:
		return node, nil
	case ast.ID:
		value, found := i.names[node.(ast.ID)]
		if !found {
			return nil, fmt.Errorf("name not found: %v", node)
		}

		return value, nil
	case ast.UnOp:
		unOp := node.(ast.UnOp)
		rawValue, err := i.Evaluate(unOp.Value)
		if err != nil {
			return nil, err
		}

		value, ok := rawValue.(ast.Integer)
		if !ok {
			return nil, fmt.Errorf("cannot use %v as integer", rawValue)
		}

		switch unOp.Op {
		case lex.UnPlus:
			return +value, nil
		case lex.UnMinus:
			return -value, nil
		}

		return nil, fmt.Errorf("interpreter: unknown unary: %s", unOp.Op)
	case ast.BinOp:
		binOp := node.(ast.BinOp)
		rawLeft, err := i.Evaluate(binOp.Left)
		if err != nil {
			return nil, err
		}
		left, ok := rawLeft.(ast.Integer)
		if !ok {
			return nil, fmt.Errorf("cannot use %v as integer", rawLeft)
		}

		rawRight, err := i.Evaluate(binOp.Right)
		if err != nil {
			return nil, err
		}
		right, ok := rawRight.(ast.Integer)
		if !ok {
			return nil, fmt.Errorf("cannot use %v as integer", rawRight)
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
			return ast.Integer(math.Pow(float64(left), float64(right))), nil
		}

		return nil, fmt.Errorf("interpreter: unknown operator: %s", binOp.Op)
	case ast.FCall:
		fcall := node.(ast.FCall)
		target, err := i.Evaluate(fcall.Target)
		if err != nil {
			return nil, err
		}

		fun, ok := target.(ast.Function)
		if !ok {
			return nil, fmt.Errorf("cannot call %s", reflect.TypeOf(target))
		}

		var args []ast.Node
		for _, arg := range fcall.Args {
			evaluated, err := i.Evaluate(arg)
			if err != nil {
				return nil, err
			}

			args = append(args, evaluated)
		}

		res, err := fun(args...)
		if err != nil {
			return nil, err
		}

		return res, nil
	case ast.FDef:
		fdef := node.(ast.FDef)
		body := func(args ...ast.Node) (ast.Node, error) {
			if len(fdef.Args) != len(args) {
				return nil, fmt.Errorf(
					"wanted %d args, got %d instead", len(fdef.Args), len(args),
				)
			}

			for index, arg := range args {
				i.names[fdef.Args[index]] = arg
			}

			return i.Evaluate(fdef.Body)
		}
		i.names[fdef.Name] = body

		return body, nil
	}

	return nil, fmt.Errorf("interpreter: unknown node: %s", node)
}
