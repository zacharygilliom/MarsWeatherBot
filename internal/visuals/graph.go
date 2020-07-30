package visuals

import (
	log "github.com/sirupsen/logrus"
	"github.com/zacharygilliom/MarsWeatherBot/internal/nasaData"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GetAxesValues() (xValues, yValues []float64) {
	data := nasaData.GetData()
	days := nasaData.GetListDays(data)
	xVals := *days
	yValues = []float64{}
	for _, day := range *days {
		temp := nasaData.GetDayTemp(day, data)
		yValues = append(yValues, temp)
	}
	xValues = IntToF64(xVals)
	return
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

func PlotGraph(xvals, yvals []float64) {
	p, err := plot.New()
	path := "/home/zacharygilliom/goProjects/MarsWeatherBot/assets/PreviousWeatherGraph.png"
	if err != nil {
		log.Debug("Error Creating New Plot: %v", err)
	}

	p.Title.Text = "Previous 5 Days Average Temperatures"
	p.X.Label.Text = "Sol Day"
	p.Y.Label.Text = "Average Temperature"
	pxys := make(plotter.XYs, len(xvals))
	for i := 0; i < len(xvals); i++ {
		pxys[i].X = xvals[i]
		pxys[i].Y = yvals[i]
	}
	err = plotutil.AddLinePoints(p,
		"Temperature", pxys,
	)
	if err != nil {
		log.Debug("Error Loading Points: %v", err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, path); err != nil {
		log.Debug("Error Saving File: %v", err)
	}

}
