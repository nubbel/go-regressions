package regressions

import "testing"

func TestLogarithmic(t *testing.T) {
	r := NewLogarithmic()
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
