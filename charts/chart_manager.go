package charts

import (
	"fmt"
	"happy_bank_simulator/models"
	"io"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type ChartsManager struct {
	page        *components.Page
	actorCharts []*ActorChart
}

// --- Instance methods ---

func (instance *ChartsManager) DrawChartsFromList() {
	instance.page = components.NewPage()
	fmt.Printf("Drawing %s charts for:\n", strconv.Itoa(len(instance.actorCharts)))
	for _, actorChart := range instance.actorCharts {
		actorChart.finalize()
		instance.page.AddCharts(actorChart.chart)
	}
	f, err := os.Create("tmp/line.html")
	if err != nil {
		panic(err)
	}
	instance.page.Render(io.MultiWriter(f))
}

func (instance *ChartsManager) FindChartForActor(actor *models.Actor) *ActorChart {
	actorName := chartNameFor(*actor)
	fmt.Printf("Looking for an existing ActorChart with name: %s\n", actorName)

	for _, actorChart := range instance.actorCharts {
		if actorChart.actorName == actorName {
			fmt.Printf("ActorChart found with chart ID %s.\n", actorChart.chart.ChartID)
			return actorChart
		}
	}

	fmt.Printf("No ActorChart found for %s.\n", actorName)
	return nil
}

func (instance *ChartsManager) CreateChartForActor(actor *models.Actor) *ActorChart {
	actorChart := ActorChart{}
	actorName := chartNameFor(*actor)
	actorChart.actorName = actorName

	lineSmoothChart := charts.NewLine()
	lineSmoothChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    actorName,
			Subtitle: "Balance at the end of simulation",
		}),
	)

	lineSmoothChart.SetSeriesOptions(charts.WithLineChartOpts(
		opts.LineChart{
			Smooth: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)

	actorChart.chart = lineSmoothChart
	fmt.Printf("New ActorChart created for %s with chart ID %s\n", actorName, actorChart.chart.ChartID)

	return &actorChart
}

func (instance *ChartsManager) AddChartToList(chart *ActorChart) {
	instance.actorCharts = append(instance.actorCharts, chart)
	fmt.Printf("ActorChart added to actorCharts list, now containing %s actorCharts.\n", strconv.Itoa(len(instance.actorCharts)))
}

// --- Package methods ---

func chartNameFor(actor models.Actor) string {
	return fmt.Sprintf("%s#%s", actor.Type, strconv.Itoa(int(actor.GetID())))
}
