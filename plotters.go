package mmio

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
		if x[i] == 0. && y[i] == 0. {
			continue
		}
		xn = append(xn, x[i])
		yn = append(yn, y[i])
	}
	if err := plotutil.AddScatters(p, points(xn, yn)); err != nil {
		log.Fatalf(" plotters.Scatter1 error: %v", err)
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

// Line creates a generic line plot
func Line(fp string, x []float64, ys map[string][]float64, width float64) {
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
	if err := p.Save(vg.Length(width)*vg.Inch, 8*vg.Inch, fp); err != nil {
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

func points(x, y []float64) plotter.XYs {
	if len(x) != len(y) {
		panic("mmplt.scatter error: unequal points array sizes")
	}
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
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

// rev is quick function used to reverse order of a slice
func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// revF is quick function used to reverse order of a float64 slice
func revF(s []float64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// onlyPositive removes all value <= 0.0 and all NaN's
func onlyPositive(s []float64) []float64 {
	var x []int
	for i := range s {
		if s[i] <= 0 || math.IsNaN(s[i]) {
			x = append(x, i)
		}
	}
	rev(x)
	for _, i := range x {
		s = append(s[:i], s[i+1:]...)
	}
	return s[:len(s)]
}