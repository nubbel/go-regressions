package regressions

import "testing"

func TestPower(t *testing.T) {
	r := NewPower()
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
