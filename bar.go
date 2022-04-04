package mmplt

import (
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Bar1 crates a simple 1-parameter bar chart
func Bar1(fp string) {
	groupA := plotter.Values{20, 35, 30, 35, 27}
	groupB := plotter.Values{25, 32, 34, 20, 25}
	groupC := plotter.Values{12, 28, 15, 21, 8}

	p := plot.New()
	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		log.Fatalf(" plotters.Bar1 error: %v", err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	barsB, err := plotter.NewBarChart(groupB, w)
	if err != nil {
		log.Fatalf(" plotters.Bar1 error: %v", err)
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(1)

	barsC, err := plotter.NewBarChart(groupC, w)
	if err != nil {
		log.Fatalf(" plotters.Bar1 error: %v", err)
	}
	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(2)
	barsC.Offset = w

	p.Add(barsA, barsB, barsC)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Bar1 error: %v", err)
	}
}

// Bar create a generic bar chart
func Bar(fp string, y []float64, xlab []string) {
	p := plot.New()
	p.Title.Text = fp

	v := make(plotter.Values, len(y))
	for i, d := range y {
		v[i] = d
	}
	p.Y.Label.Text = "Score"

	w := vg.Points(15)

	bars, err := plotter.NewBarChart(v, w)
	if err != nil {
		log.Fatalf(" plotters.Bar error: %v", err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	p.Add(bars)
	p.NominalX(xlab...)
	p.X.Tick.Label.Rotation += math.Pi / 2.
	p.X.Tick.Label.YAlign = draw.YCenter
	p.X.Tick.Label.XAlign = draw.XRight

	if err := p.Save(10*vg.Inch, 6*vg.Inch, fp); err != nil {
		log.Fatalf(" plotters.Bar error: %v", err)
	}
}
