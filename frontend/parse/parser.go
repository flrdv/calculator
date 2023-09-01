package parse

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse/ast"
	"fmt"
	"strconv"
)

type Parser struct {
	lexer *lex.Lexer
}

func NewParser(lexer *lex.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() (program ast.Program, err error) {
	for !p.lexer.EOF() {
		node, err := p.stmt()
		if err != nil {
			return nil, err
		}

		program = append(program, node)
	}

	return program, nil
}

func (p *Parser) stmt() (ast.Node, error) {
	expr, err := p.expr()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.ChFlow:
			base, ok := expr.(ast.FCall)
			if !ok {
				return nil, fmt.Errorf("cannot use %v as a function signature", expr)
			}

			name, ok := base.Target.(ast.ID)
			if !ok {
				return nil, fmt.Errorf("cannot use %v as a function name", base.Target)
			}

			fdef := ast.FDef{Name: name}

			for _, arg := range base.Args {
				id, ok := arg.(ast.ID)
				if !ok {
					return nil, fmt.Errorf("cannot use %v as a function argument", arg)
				}

				fdef.Args = append(fdef.Args, id)
			}

			fdef.Body, err = p.stmt()

			return fdef, err
		case lex.OpPlus, lex.OpMinus:
			right, err := p.expr()
			if err != nil {
				return nil, err
			}

			expr = ast.BinOp{
				Op:    lexeme.Type,
				Left:  expr,
				Right: right,
			}
		default:
			if lexeme.Type.IsSymbol() {
				return nil, fmt.Errorf("unexpected operator: %s", lexeme)
			}

			p.lexer.Back()

			return expr, nil
		}
	}

	return expr, nil
}

func (p *Parser) expr() (ast.Node, error) {
	factor, err := p.power()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.OpStar, lex.OpSlash:
			right, err := p.power()
			if err != nil {
				return nil, err
			}

			factor = ast.BinOp{
				Op:    lexeme.Type,
				Left:  factor,
				Right: right,
			}
		default:
			p.lexer.Back()

			return factor, nil
		}
	}

	return factor, nil
}

func (p *Parser) power() (ast.Node, error) {
	term, err := p.term()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.OpCaret:
			right, err := p.power()
			if err != nil {
				return nil, err
			}

			term = ast.BinOp{
				Op:    lex.OpCaret,
				Left:  term,
				Right: right,
			}
		default:
			p.lexer.Back()

			return term, nil
		}
	}

	return term, nil
}

func (p *Parser) term() (ast.Node, error) {
	factor, err := p.factor()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.LParen:
			factor, err = p.fcall(factor)
			if err != nil {
				return nil, err
			}
		default:
			p.lexer.Back()

			return factor, nil
		}
	}

	return factor, nil
}

func (p *Parser) factor() (ast.Node, error) {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return nil, err
	}

	switch lexeme.Type {
	case lex.Number:
		return strconv.ParseInt(lexeme.Value, 10, 64)
	case lex.Id:
		return lexeme.Value, nil
	case lex.UnPlus, lex.UnMinus:
		value, err := p.power()
		return ast.UnOp{
			Op:    lexeme.Type,
			Value: value,
		}, err
	case lex.LParen:
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}

		return stmt, p.match(lex.RParen)
	default:
		return nil, fmt.Errorf("unexpected factor: %s", lexeme)
	}
}

func (p *Parser) fcall(base ast.Node) (ast.Node, error) {
	var args []ast.Node

	for {
		if err := p.match(lex.RParen); err == nil {
			return ast.FCall{
				Target: base,
				Args:   args,
			}, nil
		}

		p.lexer.Back()
		arg, err := p.stmt()
		if err != nil {
			return nil, err
		}

		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		switch lexeme.Type {
		case lex.ChComma:
		case lex.RParen:
			return ast.FCall{
				Target: base,
				Args:   args,
			}, nil
		default:
			return nil, fmt.Errorf("unexpected symbol: %s (expected ) or ,)", lexeme)
		}
	}
}

func (p *Parser) match(typ lex.LexemeType) error {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return err
	}

	if lexeme.Type != typ {
		return fmt.Errorf("wanted %s, got %s", typ, lexeme)
	}

	return nil
}
