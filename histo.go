package mmplt

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Histo creates a generic histogram
func Histo(fp string, x []float64, nbins int) {
	p := plot.New()

	p.Title.Text = fp

	v := make(plotter.Values, len(x))
	for i, d := range x {
		v[i] = d
	}

	h, err := plotter.NewHist(v, nbins)
	if err != nil {
		log.Fatalf(" plotters.Histo error: %v", err)
	}

	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Histo error: %v", err)
	}
}

// HistoGT0 creates a generic histogram of all values >0.
func HistoGT0(fp string, x []float64, nbins int) {
	p := plot.New()

	n0 := 0
	for _, d := range x {
		if d <= 0. {
			n0++
		}
	}

	p.Title.Text = fmt.Sprintf("%s (n= %d; n0=%d)", fp, len(x), n0)

	v, i := make(plotter.Values, len(x)-n0), 0
	for _, d := range x {
		if d > 0. {
			v[i] = d
			i++
		}
	}

	h, err := plotter.NewHist(v, nbins)
	if err != nil {
		log.Fatalf(" plotters.HistoGT0 error: %v", err)
	}

	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.HistoGT0 error: %v", err)
	}
}
