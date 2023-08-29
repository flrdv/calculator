package operator

import "strings"

type Operator string

var (
	All              = []Operator{Plus, Minus, Star, Slash, Caret}
	RightAssociative = []Operator{Caret}
)

const (
	Plus  Operator = "+"
	Minus Operator = "-"
	Star  Operator = "*"
	Slash Operator = "/"
	Caret Operator = "^"
)

func IsRightAssociative(op Operator) bool {
	for _, operator := range RightAssociative {
		if operator == op {
			return true
		}
	}

	return false
}

func IsPrefix(opPrefix string) bool {
	for _, operator := range All {
		if strings.HasPrefix(string(operator), opPrefix) {
			return true
		}
	}

	return false
}
