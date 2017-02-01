package regressions

import "testing"

func TestLinear(t *testing.T) {
	r := NewLinear()
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
