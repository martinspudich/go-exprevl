package exprevl

import "errors"

var (
	ErrArgNotFound = errors.New("arg: not found")
	ErrExprInvalid = errors.New("expression: invalid expression")
)
