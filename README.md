# calculator

Simple calculator in Go, using top-down recursive parser. At the moment, supported:
- Basic arithmetic (including power)
- Variables
- Function calls
- Function defining

Soon:
- Namespaces
- Types
- LLVM backend (using simple interpreter at the moment)

## How to use?
```bash
git clone https://github.com/fakefloordiv/calculator
cd calculator
go run cmd/main.go
```

This will run an interactive shell. For running, go>=1.20 is required
