package main

import (
	"embed"
	"log/slog"
	"strconv"
	"strings"
)

//var grid [1000][1000]int

// 568658 : too low
// 641480 : too high
// 640376 : too high

type opRange struct {
	startX int
	startY int
	endX   int
	endY   int
}

type myGrid struct {
	grid [][]int
}

type myInterface interface {
	Op(int, opRange)
	Toggle(opRange)
	Count() int
}

func (g *myGrid) Op(operation int, lRange opRange) {

	var count int

	for i := lRange.startX; i <= lRange.endX; i++ {
		for j := lRange.startY; j <= lRange.endY; j++ {
			if g.grid[j][i] != operation {
				g.grid[j][i] = operation
				count++
			}

		}
	}

	slog.Info("DEBUG: lights", slog.Int("updated", count))

}

func (g *myGrid) Toggle(lRange opRange) {

	var count int

	for i := lRange.startX; i <= lRange.endX; i++ {
		for j := lRange.startY; j <= lRange.endY; j++ {
			count++
			if g.grid[j][i] == 0 {
				g.grid[j][i] = 1
			} else {
				g.grid[j][i] = 0
			}
		}
	}

	slog.Info("DEBUG: lights", slog.Int("toggled", count))
}

func getRanges(fields []string) opRange {

	var startX, endX, startY, endY int
	var sourceArray, targetArray []string

	// turn on, turn off or toggle
	if len(fields) > 4 {
		sourceArray = strings.Split(fields[2], ",")
		targetArray = strings.Split(fields[4], ",")
	} else {
		sourceArray = strings.Split(fields[1], ",")
		targetArray = strings.Split(fields[3], ",")
	}

	// build x ranges
	startX, _ = strconv.Atoi(sourceArray[0])
	startY, _ = strconv.Atoi(sourceArray[1])
	endX, _ = strconv.Atoi(targetArray[0])
	endY, _ = strconv.Atoi(targetArray[1])

	lRange := opRange{
		startX,
		startY,
		endX,
		endY,
	}

	return lRange
}

func (g *myGrid) Count() int {
	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if g.grid[y][x] == 1 {
				count++
			}
		}
	}
	return count
}

func initGrid(size int) *myGrid {
	lightGrid := make([][]int, size)
	for i := range lightGrid {
		lightGrid[i] = make([]int, size)
	}

	return &myGrid{grid: lightGrid}
}

//go:embed input.txt
var myFileContent embed.FS

func main() {

	var lightGrid myInterface
	var lRange opRange
	var operation string
	var fields []string

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		slog.Error("Error reading file:", err)
		return
	}

	// split content by line
	content := strings.Split(string(input), "\n")

	// init my grid
	lightGrid = initGrid(1000)

	for _, line := range content {

		fields = strings.Fields(line)

		if len(fields) > 0 {

			lRange = getRanges(fields)

			operation = fields[0] + fields[1]

			slog.Info("DEBUG:", slog.Int("startX", lRange.startX), slog.Int("startY", lRange.startY), slog.Int("endX", lRange.endX), slog.Int("endY", lRange.endY))

			switch operation[:6] {
			case "turnon":
				slog.Info("DEBUG: operation light on")
				lightGrid.Op(1, lRange)
			case "turnof":
				slog.Info("DEBUG: operation light off")
				lightGrid.Op(0, lRange)
			case "toggle":
				slog.Info("DEBUG: operation light toggle")
				lightGrid.Toggle(lRange)
			}
		}
	}

	// count the number of lights on
	slog.Info("lights on", slog.Int("count", lightGrid.Count()))
}
