package regressions

import "math"

// NewLogLog returns a new Regression for log-log regression.
func NewLogLog() Regression {
	return &logarithmic{
		base:    NewLogarithmic(),
		LogFunc: math.Log,
	}
}
