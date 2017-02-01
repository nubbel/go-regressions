package regressions

import "math"

// NewPower returns a new Regression for power regression.
func NewPower() Regression {
	return &logarithmic{
		base:    NewExponential(),
		LogFunc: math.Log,
	}
}
