package simulation

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

type ActorChart struct {
	actorName string
	items     []opts.LineData
	months    []string
	chart     *charts.Line
}

func (instance *ActorChart) AddItem(month string, itemValue float64) {
	instance.months = append(instance.months, month)
	fmt.Println("📈 New month", month, "added to chart", instance.actorName)
	instance.items = append(instance.items, opts.LineData{Value: itemValue})
	fmt.Println("📈 New item", itemValue, "added to chart", instance.actorName)
}

func (instance *ActorChart) Finalize() {
	instance.chart.SetXAxis(instance.months).AddSeries("Category A", instance.items)
	fmt.Println("📈 ActorChar", instance.actorName, "finilized")
}

// ------

type EchartsManager struct {
	page        *components.Page
	actorCharts []*ActorChart
}

func (instance *EchartsManager) AddChartToList(chart *ActorChart) {
	instance.actorCharts = append(instance.actorCharts, chart)
}

func (instance *EchartsManager) ListCharts() []*ActorChart {
	return instance.actorCharts
}

func (instance *EchartsManager) DrawChartsFromList() {
	for _, actorChart := range instance.actorCharts {
		instance.page.AddCharts(actorChart.chart)
	}
	f, err := os.Create("tmp/line.html")
	if err != nil {
		panic(err)
	}
	instance.page.Render(io.MultiWriter(f))
}

func (instance *EchartsManager) findOrCreateChartForActor(actor *models.Actor) *ActorChart {
	chart := instance.doesThisActorAlreadyHaveChart(actor)

	if chart != nil {
		return chart
	}

	return NewChartForActor(*actor)
}

func (instance *EchartsManager) doesThisActorAlreadyHaveChart(actor *models.Actor) *ActorChart {
	actorName := generateActorNameFor(*actor)
	for _, chart := range instance.ListCharts() {
		if chart.actorName == actorName {
			return chart
		}
	}
	return nil
}

// ---

func NewChartForActor(actor models.Actor) *ActorChart {
	actorChart := ActorChart{}
	actorName := generateActorNameFor(actor)
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
	fmt.Println("New ActorChart created for", actorName)

	return &actorChart
}

func generateActorNameFor(actor models.Actor) string {
	return fmt.Sprintf("%s#%s", actor.ModelName(), strconv.Itoa(int(actor.GetID())))
}
