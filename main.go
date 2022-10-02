package exprevl

import (
	"strconv"
	"strings"
)

const (
	Add   = "+"
	Sub   = "-"
	Multi = "*"
	Div   = "/"
)

// Evaluate function evaluate math expression to value.
func Evaluate(expr string, args map[string]float64) (float64, error) {
	// Find operands and operators in expression
	var values []string
	var operators []string
	var value string
	for _, char := range expr {
		switch string(char) {
		case Add, Sub, Multi, Div:
			values = append(values, strings.TrimSpace(value))
			operators = append(operators, string(char))
			value = ""
		default:
			value += string(char)
		}
	}
	values = append(values, strings.TrimSpace(value))

	// Replace arguments by values
	operands, err := replaceArgByValue(values, args)
	if err != nil {
		return 0, err
	}

	// Evaluate two numbers and operands
	result := operands[0]
	for i, operator := range operators {
		r, err := eval(result, operands[i+1], operator)
		if err != nil {
			return 0, err
		}
		result = r
	}
	return result, nil
}

// eval will evaluate two numbers with operand
func eval(a, b float64, operand string) (float64, error) {
	switch operand {
	case Add:
		return a + b, nil
	case Sub:
		return a - b, nil
	case Multi:
		return a * b, nil
	case Div:
		return a / b, nil
	}

	return 0, nil
}

// replaceArgByValue replace arguments in expression by value from args map
func replaceArgByValue(values []string, args map[string]float64) ([]float64, error) {
	var err error
	var operands []float64
	for _, v := range values {
		operand, ok := args[v]
		if !ok {
			operand, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return []float64{}, ErrExprInvalid
			}
		}
		operands = append(operands, operand)

	}
	return operands, nil
}
