package regressions

import "math"

type exponential struct {
	linear
	LogFunc func(float64) float64
	ExpFunc func(float64) float64
}

// NewExponential returns a new Regression for exponential regression.
func NewExponential() Regression {
	return &exponential{
		linear:  linear{},
		LogFunc: math.Log, // natural logarithm
		ExpFunc: math.Exp, // base-e exponential
	}
}

// Fit implements the Fitter interface.
func (r *exponential) Fit(dps ...DataPoint) error {
	logDps := make([]DataPoint, len(dps))

	for i, dp := range dps {
		logY := r.LogFunc(dp.GetY())
		if math.IsNaN(logY) {
			return ErrLogIsNaN
		}
		if math.IsInf(logY, 0) {
			return ErrLogIsInf
		}

		logDps[i] = dataPoint{dp.GetX(), logY}
	}

	err := r.linear.Fit(logDps...)
	if err != nil {
		return err
	}

	return nil
}

// Predict implements the Predicter interface.
func (r *exponential) Predict(x float64) (float64, error) {
	intercept := r.linear.impl.Coeff(0)
	slope := r.linear.impl.Coeff(1)

	y := r.ExpFunc(intercept) * r.ExpFunc(slope*x)
	if math.IsNaN(y) {
		return 0, ErrLogIsNaN
	}
	if math.IsInf(y, 0) {
		return 0, ErrLogIsInf
	}

	return y, nil
}