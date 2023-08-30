package parse

import (
	"calculator/frontend/lex"
)

type Program []Node

type Node any

type Number int64

type ID string

type BinOp struct {
	Op          lex.LexemeType
	Left, Right Node
}

type UnOp struct {
	Op    lex.LexemeType
	Value Node
}
