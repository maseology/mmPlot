package mmplt

import (
	"math"
	"time"

	"gonum.org/v1/plot/plotter"
)

func points(x, y []float64) plotter.XYs {
	if len(x) != len(y) {
		panic("mmplt.scatter error: unequal points array sizes")
	}
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		if !math.IsNaN(y[i]) {
			pts[i].X = x[i]
			pts[i].Y = y[i]
		} else {
			pts[i].X = x[i]
			pts[i].Y = 0.
		}
	}
	return pts
}

func datePoints(d []time.Time, y []float64) plotter.XYs {
	if len(d) != len(y) {
		panic("mmplt.scatter error: unequal points array sizes")
	}
	pts := make(plotter.XYs, len(d))
	for i := range pts {
		pts[i].X = float64(d[i].Unix())
		pts[i].Y = y[i]
	}
	return pts
}
