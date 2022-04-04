package mmplt

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"sort"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Line creates a generic line plot
func Line(fp string, x []float64, ys map[string][]float64, width, height float64) {
	p := plot.New()

	lines := make([]interface{}, 0)
	for l, y := range ys {
		lines = append(lines, l)
		lines = append(lines, points(x, y))
	}
	err := plotutil.AddLines(p, lines...)
	if err != nil {
		log.Fatalf(" plotters.Line error: %v", err)
	}
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(vg.Length(width)*vg.Inch, vg.Length(height)*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Line error: %v", err)
	}
}

// LineCol creates a generic line plot with specified colour scheme
func LineCol(fp string, x []float64, ys map[string][]float64, colours map[string]color.RGBA) {
	p := plot.New()

	// p.Title.Text = fp
	// p.X.Label.Text = ""
	// p.Y.Label.Text = ""

	for l, y := range ys {
		ps, err := plotter.NewLine(points(x, y))
		if err != nil {
			panic(err)
		}
		if _, ok := colours[l]; !ok {
			panic("colour not found")
		}
		ps.Color = colours[l]
		ps.Width = vg.Points(4) // line thickness
		p.Add(ps)
		p.Legend.Add(l, ps)
	}
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(16*vg.Inch, 8*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints creates a generic plot of one x and many y's
func LinePoints(fp string, x []float64, ys [][]float64) {
	p := plot.New()

	for i := range ys {
		err := plotutil.AddLinePoints(p, fmt.Sprintf("v%d", i+1), points(x, ys[i]))
		if err != nil {
			panic(err)
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints1 creates a generic plot of one xy set of data only
func LinePoints1(fp string, x, y []float64) {
	p := plot.New()

	err := plotutil.AddLinePoints(p, "v1", points(x, y))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints2 creates a generic plot of lines from 2 sets of xy data
func LinePoints2(fp string, x, y1, y2 []float64) {
	p := plot.New()

	err := plotutil.AddLinePoints(p,
		"v1", points(x, y1),
		"v2", points(x, y2))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// Temporal creates a generic line plot, but based on dates
func Temporal(fp string, dts []time.Time, ys map[string][]float64, width float64) {
	p := plot.New()

	lines := make([]interface{}, 0)
	for l, y := range ys {
		lines = append(lines, l)
		lines = append(lines, datePoints(dts, y))
	}
	err := plotutil.AddLinePoints(p, lines...)
	if err != nil {
		panic(err)
	}
	p.Legend.Top = true
	p.X.Tick.Marker = plot.TimeTicks{Format: "Jan06"} // "2006-01-02\n15:04"}

	// Save the plot to a PNG file.
	if err := p.Save(vg.Length(width)*vg.Inch, 8*vg.Inch, fp); err != nil {
		panic(err)
	}
}

func sequentialLine(v []float64) plotter.XYs {
	pts, c := make(plotter.XYs, len(v)), 0
	for i := range pts {
		if math.IsNaN(v[i]) {
			continue
		}
		pts[c].X = float64(i)
		pts[c].Y = v[i]
		c++
	}
	return pts[:c]
}

func cumulativeDistributionLine(v []float64) plotter.XYs {
	v = onlyPositive(v)
	sort.Float64s(v)
	revF(v)
	pts, c, x := make(plotter.XYs, len(v)), 0, float64(len(v))/100.
	for i := range pts {
		if math.IsNaN(v[i]) {
			continue
		}
		pts[c].X = float64(i) / x
		pts[c].Y = v[i]
		c++
	}
	return pts[:c]
}
