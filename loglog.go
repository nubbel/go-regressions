package regressions

import "math"

// NewLogLog returns a new Regression for log-log regression.
func NewLogLog() Regression {
	return NewLogarithmicWithLogFunc(
		math.Log, // natural logarithm
	)
}

// NewLogLogWithLogFunc returns a new Regression for log-log regression.
func NewLogLogWithLogFunc(log func(float64) float64) Regression {
	return &logarithmic{
		base:    NewLogarithmicWithLogFunc(log),
		LogFunc: log,
	}
}
