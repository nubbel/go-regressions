package regressions

// DataPoint represents a data point in the regression.
type DataPoint interface {
	GetX() float64
	GetY() float64
}

type dataPoint struct {
	x float64
	y float64
}

// GetX implements the DataPoint interface.
func (dp dataPoint) GetX() float64 {
	return dp.x
}

// GetY implements the DataPoint interface.
func (dp dataPoint) GetY() float64 {
	return dp.y
}

// Fitter is the interface that wraps the Fit method.
type Fitter interface {
	Fit(dps ...DataPoint) error
}

// Predicter is the interface that wraps the Predict method.
type Predicter interface {
	Predict(x float64) (float64, error)
}

// Regression is the interface that groups the Fit and Predict methods.
type Regression interface {
	Fitter
	Predicter

	// GetR2 returns the R-squared metric of the regression.
	GetR2() float64
}
