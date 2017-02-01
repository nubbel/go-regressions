package regressions

import "math"

type logarithmic struct {
	base    Regression
	LogFunc func(float64) float64
}

// NewLogarithmic returns a new Regression for logarithmic regression.
func NewLogarithmic() Regression {
	return &logarithmic{
		base:    NewLinear(),
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

	return r.base.Fit(logDps...)
}

// Predict implements the Predicter interface.
func (r *logarithmic) Predict(x float64) (float64, error) {
	logX := r.LogFunc(x)
	if math.IsNaN(logX) {
		return 0, ErrLogIsNaN
	}
	if math.IsInf(logX, 0) {
		return 0, ErrLogIsInf
	}

	return r.base.Predict(logX)
}

// GetR2 implements the Regression interface.
func (r *logarithmic) GetR2() float64 {
	return r.base.GetR2()
}
