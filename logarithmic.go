package regressions

import "math"

type logarithmic struct {
	linear
	LogFunc func(float64) float64
}

// NewLogarithmic returns a new Regression for logarithmic regression.
func NewLogarithmic() Regression {
	return &logarithmic{
		linear:  linear{},
		LogFunc: math.Log, // natural logarithm
	}
}

// Fit implements the Fitter interface.
func (r *logarithmic) Fit(dps ...DataPoint) error {
	logDps := make([]DataPoint, len(dps))

	for i, dp := range dps {
		logX := r.LogFunc(dp.GetX())
		if math.IsNaN(logX) {
			return ErrLogIsNaN
		}
		if math.IsInf(logX, 0) {
			return ErrLogIsInf
		}

		logDps[i] = dataPoint{logX, dp.GetY()}
	}

	return r.linear.Fit(logDps...)
}

// Predict implements the Predicter interface.
func (r *logarithmic) Predict(x float64) (float64, error) {
	return r.linear.Predict(r.LogFunc(x))
}
