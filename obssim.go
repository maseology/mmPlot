package mmplt

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// ObsSim is used to create simple observed vs. simulated hydrographs
func ObsSim(fp string, o, s []float64) {
	p := plot.New()

	// p.Title.Text = fp
	p.X.Label.Text = ""
	p.Y.Label.Text = "discharge"

	ps, err := plotter.NewLine(sequentialLine(s))
	if err != nil {
		log.Fatalf(" plotters.ObsSim error: %v", err)
	}
	ps.Color = color.RGBA{R: 255, A: 255}

	po, err := plotter.NewLine(sequentialLine(o))
	if err != nil {
		log.Fatalf(" plotters.ObsSim error: %v", err)
	}
	po.Color = color.RGBA{B: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(ps, po)
	p.Legend.Add("obs", po)
	p.Legend.Add("sim", ps)
	p.Legend.Top = true
	// p.X.Tick.Marker = plot.TimeTicks{Format: "Jan"}

	// Save the plot to a PNG file.
	if err := p.Save(24*vg.Inch, 8*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.ObsSim error: %v", err)
	}
}

// ObsSimFDC is used to create simple observed vs. simulated flow-duration curves
func ObsSimFDC(fp string, o, s []float64) {
	p := plot.New()

	// create copies
	ocopy, scopy := make([]float64, len(o)), make([]float64, len(s))
	copy(ocopy, o)
	copy(scopy, s)

	// p.Title.Text = fp
	p.X.Label.Text = ""
	p.Y.Label.Text = "discharge"

	ps, err := plotter.NewLine(cumulativeDistributionLine(scopy))
	if err != nil {
		log.Fatalf(" plotters.ObsSimFDC error: %v", err)
	}
	ps.Color = color.RGBA{R: 255, A: 255}

	po, err := plotter.NewLine(cumulativeDistributionLine(ocopy))
	if err != nil {
		log.Fatalf(" plotters.ObsSimFDC error: %v", err)
	}
	po.Color = color.RGBA{B: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(ps, po)
	p.Legend.Add("obs", po)
	p.Legend.Add("sim", ps)
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.ObsSimFDC error: %v", err)
	}
}
