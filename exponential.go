package regressions

import "math"

type exponential struct {
	base    linear
	LogFunc func(float64) float64
	ExpFunc func(float64) float64
}

// NewExponential returns a new Regression for exponential regression.
func NewExponential() Regression {
	return &exponential{
		base:    linear{},
		LogFunc: math.Log, // natural logarithm
		ExpFunc: math.Exp, // base-e exponential
	}
}

// Fit implements the Fitter interface.
func (r *exponential) Fit(dps ...DataPoint) error {
	logDps := make([]DataPoint, len(dps))

	for i, dp := range dps {
		y := dp.GetY()
		logY := r.LogFunc(y)
		if math.IsNaN(logY) || math.IsInf(logY, 0) {
			return ErrLogUndefined(y)
		}

		logDps[i] = dataPoint{dp.GetX(), logY}
	}

	err := r.base.Fit(logDps...)
	if err != nil {
		return err
	}

	return nil
}

// Predict implements the Predicter interface.
func (r *exponential) Predict(x float64) (float64, error) {
	intercept := r.base.impl.Coeff(0)
	slope := r.base.impl.Coeff(1)

	y := r.ExpFunc(intercept) * r.ExpFunc(slope*x)

	return y, nil
}

// GetR2 implements the Regression interface.
func (r *exponential) GetR2() float64 {
	return r.base.GetR2()
}
