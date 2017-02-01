package regressions

import "testing"

func TestLinear(t *testing.T) {
	r := NewLinear()
	err := r.Fit(
		dataPoint{x: 6, y: 30},
		dataPoint{x: 20, y: 12},
		dataPoint{x: 26, y: 9},
		dataPoint{x: 30, y: 8},
		dataPoint{x: 37, y: 6},
		dataPoint{x: 42, y: 5},
		dataPoint{x: 47, y: 4},
		dataPoint{x: 54, y: 3},
		dataPoint{x: 64, y: 2},
	)
	if err != nil {
		t.Fatal(err)
	}

	p, err := r.Predict(25)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Predict(4): %+v", p)
	t.Logf("R-squared: %+v", r.GetR2())
}
