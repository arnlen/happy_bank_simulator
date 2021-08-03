package charts

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
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
	fmt.Printf("📈 New month \"%s\" added to chart #%s for %s\n",
		month, instance.chart.ChartID, instance.actorName)

	instance.items = append(instance.items, opts.LineData{Value: itemValue})
	fmt.Printf("📈 New item \"%1.2f\" added to chart #%s for %s\n",
		itemValue, instance.chart.ChartID, instance.actorName)
}

func (instance *ActorChart) Print() *ActorChart {
	fmt.Println("📈 ActorChart details:")
	fmt.Printf("- ChartID: %s\n", instance.chart.ChartID)

	monthsString := "- Months:"
	for index, month := range instance.months {
		if index != 0 {
			monthsString = fmt.Sprintf("%s,", monthsString)
		}
		monthsString = fmt.Sprintf("%s %s", monthsString, month)
	}

	itemsString := "- Items:"
	for index, item := range instance.items {
		if index != 0 {
			itemsString = fmt.Sprintf("%s,", itemsString)
		}
		itemsString = fmt.Sprintf("%s %s", itemsString, item.Name)
	}

	fmt.Println(monthsString)
	fmt.Println(itemsString)

	return instance
}

func (instance *ActorChart) finalize() {
	instance.chart.SetXAxis(instance.months).AddSeries("Category A", instance.items)
	fmt.Printf("📈 Chart %s (%s) with %s items and %s months finalized.\n",
		instance.actorName,
		instance.chart.ChartID,
		strconv.Itoa(len(instance.items)),
		strconv.Itoa(len(instance.months)),
	)
}

// --- PACKAGE METHODS ---

func NewActorChartFor(actor *models.Actor) *ActorChart {
	actorName := chartNameFor(*actor)
	fmt.Printf("Creating new ActorChart for %s.\n", actorName)
	actorChart := newActorChart()
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
	fmt.Printf("New ")
	actorChart.Print()

	return actorChart
}

func newActorChart() *ActorChart {
	return &ActorChart{
		actorName: "",
		items:     []opts.LineData{},
		months:    []string{},
		chart:     &charts.Line{},
	}
}
