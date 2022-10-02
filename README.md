# Expression Evaluator

This go library is for evaluating string expressions. The library is designed to calculate value of the expression.
Currently, only simple expressions are supported and are evaluated from left to right. Parentheses are not supported.

It contains only one public method.

```go
func Evaluate(expr string, map[string]float64)
```
Evaluate method will evaluate an expression with arguments defined in the map.

Example:

```go
import eevl github.com/martinspudich/go-exprevl

args := map[string]float64 {
  "x": 1,
  "y": 2,
}
result := eevl.evaluate("x + y, args)
println("result: ", result) // prints "result: 3"
```

## Supported math operators

For now, just basic math operators are supported.

Supported operators:

| Operator | Description |
|----------|-------------|
| +        | plus        |
| -        | minus       |
| *        | multiply    |
| /        | devision    |