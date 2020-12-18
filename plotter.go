package bandwidth

//go:generate go run main.go

import (
	"os"
	"strings"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

func collectData(s []bwData) ([]float64, []float64, []time.Time) {
	transmitted := make([]float64, 0)
	received := make([]float64, 0)
	timestamps := make([]time.Time, 0)
	for _, data := range s {
		transmitted = append(transmitted, data.tx)
		received = append(received, data.rx)
		timestamps = append(timestamps, data.timestamp)
	}
	return transmitted, received, timestamps

}

func Plot() {
	tx, rx, t := collectData(Stats)
	graphRX := chart.Chart{
		Title: "Received bytes",
		XAxis: chart.XAxis{
			Name: "Time Frame",
		},
		YAxis: chart.YAxis{
			Name: "GBs",
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				YValues: rx,
				XValues: t,
			},
		},
	}

	graphTX := chart.Chart{
		Title: "Transmitted bytes",
		XAxis: chart.XAxis{
			Name: "Time Frame",
		},
		YAxis: chart.YAxis{
			Name: "GBs",
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				YValues: tx,
				XValues: t,
			},
		},
	}

	f, _ := os.Create("output/rx" + strings.Split(time.Now().String(), ".")[0] + ".png")
	defer f.Close()
	graphRX.Render(chart.PNG, f)

	f2, _ := os.Create("output/tx" + strings.Split(time.Now().String(), ".")[0] + ".png")
	defer f2.Close()
	graphTX.Render(chart.PNG, f2)

}
