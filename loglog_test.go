package regressions

import (
	"math"
	"testing"
)

func TestLogLog(t *testing.T) {
	r := NewLogLogWithLogFunc(math.Log1p)
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
