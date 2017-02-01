package regressions

import "testing"

func TestLogLog(t *testing.T) {
	r := NewLogLog()
	err := r.Fit(testDataPoints...)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("R-squared: %+v", r.GetR2())
	evaluateRegression(t, r)
}
