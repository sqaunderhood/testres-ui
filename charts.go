package main

import (
	"fmt"
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

var (
	ColorRed   = drawing.ColorFromHex("DC1E10")
	ColorGreen = drawing.ColorFromHex("038832")
	ColorGrey  = drawing.ColorFromHex("C7C5CA")

	StyleRed = chart.Style{
		Show:        true,
		FillColor:   ColorRed,
		StrokeColor: ColorRed.WithAlpha(64),
	}

	StyleGreen = chart.Style{
		Show:        true,
		FillColor:   ColorGreen,
		StrokeColor: ColorGreen.WithAlpha(64),
	}

	StyleGrey = chart.Style{
		Show:        true,
		FillColor:   ColorGrey,
		StrokeColor: ColorGrey.WithAlpha(64),
	}
)

func drawTwoAxesChart(res http.ResponseWriter, req *http.Request) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{Show: true},
		},

		YAxis: chart.YAxis{
			Style: chart.Style{Show: true},
		},

		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: ColorGreen,
					FillColor:   ColorGreen,
				},
				Name:    "Passed",
				XValues: []float64{1.0, 3.0, 5.0, 8.0, 10.0},
				YValues: []float64{8.0, 6.0, 7.0, 10.0, 15.0},
			},

			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: ColorRed,
					FillColor:   ColorRed,
				},
				Name:    "Failed",
				XValues: []float64{1.0, 3.0, 5.0, 8.0, 10.0},
				YValues: []float64{6.0, 5.0, 4.0, 6.0, 14.0},
			},

			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: ColorGrey.WithAlpha(64),
					FillColor:   ColorGrey,
				},
				Name:    "Unknown",
				XValues: []float64{1.0, 3.0, 5.0, 8.0, 10.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 3.0},
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := graph.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering two axes chart: %v\n", err)
	}
}

func drawBarChart(res http.ResponseWriter, req *http.Request) {

	sbc := chart.BarChart{
		Title:      "Bar Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:     200,
		BarSpacing: 1,
		BarWidth:   15,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{
			{Value: 5, Style: StyleGreen},
			{Value: 6, Style: StyleRed},
			{Value: 8, Style: StyleRed},
			{Value: 3, Style: StyleRed},
			{Value: 5, Style: StyleGreen},
			{Value: 5, Style: StyleGreen},
			{Value: 5, Style: StyleGreen},
			{Value: 4, Style: StyleRed},
			{Value: 3, Style: StyleRed},
			{Value: 5, Style: StyleGreen},
			{Value: 3, Style: StyleRed},
			{Value: 9, Style: StyleGreen},
			{Value: 5, Style: StyleGreen},
			{Value: 9, Style: StyleGreen},
			{Value: 5, Style: StyleGreen},
			{Value: 5, Style: StyleRed},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := sbc.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering bar chart: %v\n", err)
	}
}

func drawStackedBarChart(res http.ResponseWriter, req *http.Request) {

	sbc := chart.StackedBarChart{
		Title:      "Bar Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:     200,
		BarSpacing: 1,
		Bars: []chart.StackedBar{
			{
				Name: "Bar 1",
				Values: []chart.Value{
					{Value: 15, Label: "Blue", Style: StyleRed},
					{Value: 7, Label: "Green", Style: StyleGreen},
					{Value: 8, Label: "Gray", Style: StyleGrey},
				},
				Width: 10,
			},
			{
				Name: "Bar 2",
				Values: []chart.Value{
					{Value: 9, Label: "Blue", Style: StyleRed},
					{Value: 6, Label: "Green", Style: StyleGreen},
					{Value: 3, Label: "Gray", Style: StyleGrey},
				},
				Width: 10,
			},
			{
				Name: "Bar 3",
				Values: []chart.Value{
					{Value: 6, Label: "Blue", Style: StyleRed},
					{Value: 10, Label: "Green", Style: StyleGreen},
					{Value: 2, Label: "Gray", Style: StyleGrey},
				},
				Width: 10,
			},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := sbc.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering bar chart: %v\n", err)
	}
}

func drawPieChart(res http.ResponseWriter, req *http.Request) {

	pie := chart.PieChart{
		Width:  250,
		Height: 250,
		Values: []chart.Value{
			{Value: 4, Label: "Failed", Style: StyleRed},
			{Value: 3, Label: "Passed", Style: StyleGreen},
			{Value: 3, Label: "Unknown", Style: StyleGrey},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := pie.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}
