package mmplt

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Wbal used to review waterbalance
func Wbal(fp string, f, a, q, g, s []float64) {
	p := plot.New()

	err := plotutil.AddLines(p,
		"pre", sequentialLine(f),
		"aet", sequentialLine(a),
		"ro", sequentialLine(q),
		"rch", sequentialLine(g),
		"sto", sequentialLine(s))
	if err != nil {
		log.Fatalf(" plotters.Wbal error: %v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Wbal error: %v", err)
	}
}
