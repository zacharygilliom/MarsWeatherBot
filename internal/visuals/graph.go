package visuals

import (
	log "github.com/sirupsen/logrus"
	chart "github.com/wcharczuk/go-chart"
	"github.com/zacharygilliom/MarsWeatherBot/internal/nasaData"
	"os"
)

func GraphWeather() {
	data := nasaData.GetData()
	days := nasaData.GetListDays(data)
	xVals := *days
	yValues := []float64{}
	for _, day := range *days {
		temp := nasaData.GetDayTemp(day, data)
		yValues = append(yValues, temp)
	}
	xValues := IntToF64(xVals)
	filename := "/home/zacharygilliom/goProjects/MarsWeatherBot/assets/PreviousWeatherGraph.png"
	f, err := os.Create(filename)
	if err != nil {
		log.Debug("Failed to Create file: %v: %v", filename, err)
	}
	defer f.Close()
	graph := chart.Chart{
		Title:      "Previous 5 Days Average Temperature",
		TitleStyle: getTitleStyle(),
		XAxis: chart.XAxis{
			Name: "Sol Day",
		},
		YAxis: chart.YAxis{
			Name: "Average Temperature",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
			},
		},
	}
	err = graph.Render(chart.PNG, f)
	if err != nil {
		log.Debug("Unable to render graph: %v", err)
	}
}

func IntToF64(ar []int) []float64 {
	newar := make([]float64, len(ar))
	var v int
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}

func getTitleStyle() chart.Style {
	return chart.Style{
		Show: true,
	}
}
