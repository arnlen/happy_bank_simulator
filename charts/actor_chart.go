package charts

import (
	"fmt"
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
	fmt.Println("📈 New month", month, "added to chart", instance.actorName)
	instance.items = append(instance.items, opts.LineData{Value: itemValue})
	fmt.Println("📈 New item", itemValue, "added to chart", instance.actorName)
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
