package charts

import (
	"fmt"
	"happy_bank_simulator/models"
	"io"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
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

func (instance *ChartsManager) FindActorChartInListFor(actor *models.Actor) *ActorChart {
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

func (instance *ChartsManager) AddChartToList(chart *ActorChart) {
	instance.actorCharts = append(instance.actorCharts, chart)
	fmt.Printf("ActorChart added to actorCharts list, now containing %s actorCharts.\n",
		strconv.Itoa(len(instance.actorCharts)))
}

// --- Package methods ---

func chartNameFor(actor models.Actor) string {
	return fmt.Sprintf("%s#%s", actor.Type, strconv.Itoa(int(actor.GetID())))
}
