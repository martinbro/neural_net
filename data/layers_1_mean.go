package data

import (
	"gonum.org/v1/gonum/mat"
)

func Layers_1_mean() *mat.Dense {
	return mat.NewDense(5, 1, []float64{
		0.71084273,
		0.75236458,
		0.67221397,
		0.71036267,
		0.75235319,
	})
}
