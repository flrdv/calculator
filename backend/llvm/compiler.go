package llvm

import (
	"github.com/llir/llvm/ir/value"
)

type Compiler struct {
	names map[string]value.Value
}

func NewCompiler(names map[string]value.Value) Compiler {
	if names == nil {
		names = map[string]value.Value{}
	}

	return Compiler{
		names: names,
	}
}

func (c Compiler) Compile() (string, error) {
	return "", nil
}
