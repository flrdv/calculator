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

### Syntax
Enter an expression, the result will be printed on the next line.

#### Define function
```
f(x, y) -> x + y
```

Note: function body is always a single expression. The result of it is returned as a return value

#### Call function
```
f(x, y)
```

#### Define variable
```
x -> 5
```

Note: variables defining is an expression, returning value of a newly defined variable. By that, we also can use the following form:
```
x -> y -> 5
```
...resulting in x == y == 5

#### Binary operations
`+` - add
`-` - subtract
`/` - divide
`*` - multiply
`^` - power

#### Unary operations
`+` and `-` respectively. Note: the precedence if unary operations are higher than any math operation, except power and function calls
