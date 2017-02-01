package regressions

import "testing"

var testDataPoints = []DataPoint{
	dataPoint{x: 0.85, y: 91},
	dataPoint{x: 1, y: 78},
	dataPoint{x: 2, y: 64},
	dataPoint{x: 3, y: 54},
	dataPoint{x: 4, y: 47},
	dataPoint{x: 5, y: 42},
	dataPoint{x: 6, y: 37},
	dataPoint{x: 8, y: 30},
	dataPoint{x: 9, y: 26},
	dataPoint{x: 10, y: 25},
	dataPoint{x: 12, y: 20},
	dataPoint{x: 30, y: 6},
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
