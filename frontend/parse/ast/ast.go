package ast

import (
	"calculator/frontend/lex"
)

type Program []Node

type (
	Node    any
	Integer = int64
	Float   = float64
	ID      = string
	BinOp   struct {
		Op          lex.LexemeType
		Left, Right Node
	}
	UnOp struct {
		Op    lex.LexemeType
		Value Node
	}
	Function = func(...Node) (Node, error)
	FCall    struct {
		Target Node
		Args   []Node
	}
	FDef struct {
		Name string
		Args []string
		Body Node
	}
)
