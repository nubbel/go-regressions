package regressions

import "github.com/sajari/regression"

type linear struct {
	impl regression.Regression
}

// NewLinear returns a new Regression for linear regression.
func NewLinear() Regression {
	return &linear{}
}

// Fit implements the Fitter interface.
func (r *linear) Fit(dps ...DataPoint) error {
	for _, dp := range dps {
		r.impl.Train(regression.DataPoint(dp.GetY(), []float64{dp.GetX()}))
	}

	return r.impl.Run()
}

// Predict implements the Predicter interface.
func (r *linear) Predict(x float64) (float64, error) {
	return r.impl.Predict([]float64{x})
}

// GetR2 implements the Regression interface.
func (r *linear) GetR2() float64 {
	return r.impl.R2
}

func (r *linear) String() string {
	return r.impl.String()
}
