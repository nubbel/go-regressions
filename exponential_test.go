package regressions

import "testing"

func TestExponential(t *testing.T) {
	r := NewExponential()
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
