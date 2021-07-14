package charts

import "github.com/go-echarts/go-echarts/v2/opts"

type node struct {
	ItemStyle  map[string]string `json:"itemStyle"`
	Name       string            `json:"name"`
	Y          float64           `json:"y"`
	X          float64           `json:"x"`
	SymbolSize float64           `json:"symbolSize"`
}

type link struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type graphData struct {
	Nodes []node `json:"nodes"`
	Links []link `json:"links"`
}

type Data struct {
	Nodes []opts.GraphNode
	Links []opts.GraphLink
}