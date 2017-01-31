package regressions

import "testing"

func TestLogarithmic(t *testing.T) {
	r := NewLogarithmic()
	err := r.Fit(
		dataPoint{x: 1, y: 1},
		dataPoint{x: 2, y: 4},
		dataPoint{x: 3, y: 9},
	)
	if err != nil {
		t.Fatal(err)
	}

	p, err := r.Predict(4)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Predict(4): %+v", p)
	t.Logf("R-squared: %+v", r.GetR2())
}
