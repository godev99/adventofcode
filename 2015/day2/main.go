package main

import (
	//"bufio"
	"embed"
	"fmt"
	"strconv"
	"strings"
	//"os"
)

//go:embed input.txt
var myFileContent embed.FS

type Present struct {
	height int
	width  int
	length int
}

func NewPresent(l int, w int, h int) *Present {
	return &Present{length: l, width: w, height: h}
}

// ribon returns the ribon of the present
func (p *Present) Ribon() int {
	var greatest int
	area := []int{}
	for _, side := range []int{p.height, p.width, p.length} {
		if side > greatest {
			if greatest != 0 {
				area = append([]int{greatest}, area...)
			}
			greatest = side

		} else {
			area = append(area, side)
		}
	}

	ribon := 2*area[0] + 2*area[1] + p.height*p.width*p.length

	return ribon
}

// squareFeet returns the square feet of the present
func (p *Present) WrappingPaper() int {
	areas := []int{p.height * p.width, p.length * p.width, p.height * p.length}
	smallest := areas[0]
	for _, area := range areas {
		if area < smallest {
			smallest = area
		}
	}
	squareFeet := areas[0]*2 + areas[1]*2 + areas[2]*2 + smallest

	return squareFeet
}

func main() {

	fmt.Println("Day 2: I Was Told There Would Be No Math")
	fmt.Println("")

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// split content by line
	content := strings.Split(string(input), "\n")

	// square feets
	var sf int

	// ribon
	var ribon int

	for _, line := range content {
		if line != "" {
			tmp := strings.Split(line, "x")

			length, err := strconv.Atoi(tmp[0])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			width, err := strconv.Atoi(tmp[1])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			height, err := strconv.Atoi(tmp[2])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			//newPresent := NewPresent(length, width, height)
			// create a new present with NewPresent constructor
			newPresent := NewPresent(length, width, height)
			sf = sf + newPresent.WrappingPaper()
			ribon = ribon + newPresent.Ribon()
			fmt.Println("Present:", newPresent, "Wrapping paper:", newPresent.WrappingPaper(), "Ribon:", newPresent.Ribon())
		}
	}

	fmt.Println("")
	fmt.Println("Wrapping paper:", sf)
	fmt.Println("Ribon:", ribon)

}
