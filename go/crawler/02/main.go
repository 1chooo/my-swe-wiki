package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart"
)

func main() {
	csvFile, err := os.Open("../csv/number_of_films.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	x := make([]string, 0)
	y := make([]int, 0)

	for _, row := range rows {
		x = append(x, row[0])
		numStr := strings.ReplaceAll(row[1], ",", "")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}
		y = append(y, num)
	}

	fmt.Println(x, y)

	graph := chart.BarChart{
		Title:      "Number of Films",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		XAxis: chart.Style{
			TextRotationDegrees: 45,
		},
		Series: []chart.Series{
			chart.BarSeries{
				Name: "Films",
				Values: y,
			},
		},
	}

	f, _ := os.Create("chart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
