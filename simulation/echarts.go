package simulation

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntLine = 6
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Smooth Style mixed ((title and label options)",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		lineSmooth(),
	)
	f, err := os.Create("tmp/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
