package regressions

import "testing"

var testDataPoints = []DataPoint{
	dataPoint{x: 6, y: 30},
	dataPoint{x: 20, y: 12},
	dataPoint{x: 25, y: 10},
	dataPoint{x: 26, y: 9},
	dataPoint{x: 30, y: 8},
	dataPoint{x: 37, y: 6},
	dataPoint{x: 42, y: 5},
	dataPoint{x: 47, y: 4},
	dataPoint{x: 54, y: 3},
	dataPoint{x: 64, y: 2},
}

func evaluateRegression(t *testing.T, r Regression) {
	for _, dp := range testDataPoints {
		x := dp.GetX()
		observed := dp.GetY()
		predicted, err := r.Predict(x)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%v:\t%v (observed)\t%v (predicted)", x, observed, predicted)
	}
}
