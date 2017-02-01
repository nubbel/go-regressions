package regressions

import "math"

// NewPower returns a new Regression for power regression.
func NewPower() Regression {
	return NewPowerWithLogFunc(
		math.Log, // natural logarithm
	)
}

// NewPowerWithLogFunc returns a new Regression for power regression.
func NewPowerWithLogFunc(log func(float64) float64) Regression {
	return &logarithmic{
		base:    NewExponential(),
		LogFunc: log,
	}
}
