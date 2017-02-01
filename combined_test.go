package regressions

import (
	"math"
	"testing"
)

func TestMaxCombinedLogExp(t *testing.T) {
	r := NewCombined(math.Max, NewLogarithmic(), NewExponential())
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
