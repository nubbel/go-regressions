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

func TestMeanCombinedLogLogExp(t *testing.T) {
	mean := func(x, y float64) float64 {
		return (x + y) / 2
	}

	r := NewCombined(mean, NewLogLog(), NewExponential())
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
