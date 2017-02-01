package regressions

import "fmt"

// ErrLogUndefined signifies that the logarithm of the underlying value is undefined.
type ErrLogUndefined float64

// Value returns the floating point value for which the logarithm is undefined.
func (f ErrLogUndefined) Value() float64 {
	return float64(f)
}

// Error implements the error interface
func (f ErrLogUndefined) Error() string {
	return fmt.Sprintf("Logarithm of %g is undefined.", f.Value())
}
