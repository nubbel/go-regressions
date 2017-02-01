package regressions

import "fmt"

// ErrLogUndefined signifies that the logarithm of the underlying value is undefined.
type ErrLogUndefined float64

func (f ErrLogUndefined) Error() string {
	return fmt.Sprintf("Logarithm of %g is undefined.", float64(f))
}
