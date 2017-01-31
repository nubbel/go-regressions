package regressions

import "errors"

var (
	ErrLogIsNaN = errors.New("Logarithm is undefined")
	ErrLogIsInf = errors.New("Logarithm is Infinity")
)
