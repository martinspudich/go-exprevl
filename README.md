# Expression Evaluator

This GO library is for evaluating string expressions.

Library is design to calculate expression to the value.

It will contain only one public interface ExpressionCalculator

```go
type ExpressionCalculator interface {
  Calculate(expr string, map[string]float64)
}
```

or only one public function

```go
func (expr string, map[string]float64)
```

Calculate method will evalueate expression with arguments deffined in map.

Example with interface

```go
ex := NewExpressionCalculator()
args := map[string]float64 {
  "x": 1,
  "y": 2,
}
result := ex.Calculate("$x + $y, args)
println("result: ", result) // prints "result: 3"
```

Example with func

```go
args := map[string]float64 {
  "x": 1,
  "y": 2,
}
result := goexprcalc.Calculate("$x + $y, args)
println("result: ", result) // prints "result: 3"
```