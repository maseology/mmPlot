package mmplt

import (
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Scatter creates a generic scatter plot
func Scatter(fp string, x, y []float64) {
	p := plot.New()

	p.Title.Text = fp
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddScatters(p, points(x, y))
	if err != nil {
		log.Fatalf(" plotters.Scatter error: %v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Scatter error: %v", err)
	}
}

// Scatter11 creates a generic scatter plot, with a 1:1 line
func Scatter11(fp string, x, y []float64) {
	p := plot.New()

	p.Title.Text = fp
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	xn, yn := []float64{}, []float64{}
	for i := range x {
		if math.IsNaN(x[i]) || math.IsNaN(y[i]) {
			continue
		}
		if x[i] == 0. && y[i] == 0. {
			continue
		}
		xn = append(xn, x[i])
		yn = append(yn, y[i])
	}
	if err := plotutil.AddScatters(p, points(xn, yn)); err != nil {
		log.Fatalf(" plotters.Scatter11 error: %v", err)
	}
	max, min := math.Max(p.X.Max, p.Y.Max), math.Min(p.X.Min, p.Y.Min)
	p.X.Max = max
	p.Y.Max = max
	p.X.Min = min
	p.Y.Min = min
	abline, iabline := make(plotter.XYs, 2), make([]interface{}, 1)
	abline[0].X, abline[0].Y = min, min
	abline[1].X, abline[1].Y = max, max
	iabline[0] = abline
	if err := plotutil.AddLines(p, iabline...); err != nil {
		log.Fatalf(" plotters.Scatter1 error: %v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Scatter1 error: %v", err)
	}
}
