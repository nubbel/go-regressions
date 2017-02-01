package regressions

type combined struct {
	combiner func(x, y float64) float64
	first    Regression
	second   Regression
}

// NewCombined returns a new Regression combining other regression.
func NewCombined(combiner func(x, y float64) float64, first, second Regression) Regression {
	return &combined{
		combiner,
		first,
		second,
	}
}

// Fit implements the Fitter interface.
func (r *combined) Fit(dps ...DataPoint) error {
	err := r.first.Fit(dps...)
	if err != nil {
		return err
	}

	return r.second.Fit(dps...)
}

// Predict implements the Predicter interface.
func (r *combined) Predict(x float64) (float64, error) {
	firstY, err := r.first.Predict(x)
	if err != nil {
		return 0, err
	}

	secondY, err := r.second.Predict(x)
	if err != nil {
		return 0, err
	}

	return r.combiner(firstY, secondY), nil
}

// GetR2 implements the Regression interface.
func (r *combined) GetR2() float64 {
	return (r.first.GetR2() + r.second.GetR2()) / 2
}
