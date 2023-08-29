package parse

import "calculator/frontend/operator"

type Program []Node

type Node any

type Number int64

type ID string

type BinOp struct {
	Op          operator.Operator
	Left, Right Node
}
