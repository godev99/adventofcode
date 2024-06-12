package main

import (
	"fmt"
	"reflect"
	"embed"
)

//go:embed input.txt
var myFileContent embed.FS

func main() {
	fmt.Println("Day 3: Perfectly Spherical Houses in a Vacuum")
	fmt.Println("")

	var absX, absY int
	var isAlreadyVisited bool

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// init matrix with input length
	fmt.Println("input length: ", len(input))
	length := len(input)
	var matrix = make([][]int, length+1)

	// first mouvement is a present for first leaving position
	matrix[0] = []int{absX, absY}
	presents := 1
	fmt.Println("first present distributed at:", absX, absY)

	//parse input
	for index, c := range input {
		switch string(c) {
			case "^":
				fmt.Println("up")
				absY++
			case "v":
				fmt.Println("down")
				absY--
			case ">":
				fmt.Println("right")
				absX++
			case "<":
				fmt.Println("left")
				absX--
			default:
		}

		// check if the house is already visited
		// if not, add it to the matrix
		isAlreadyVisited = alreadyVisited(matrix, []int{absX, absY})
		if !isAlreadyVisited {
			presents++
			fmt.Println("new present distributed at:", absX, absY)
		}

		matrix[index+1] = []int{absX, absY}
	}

	fmt.Println("")
	fmt.Println("presents distributed:", presents)
}

// check if the house is already visited
func alreadyVisited(matrix [][]int, current []int) bool {
	//fmt.Println("matrix:", matrix, "current:", current)
	for _, house := range matrix {
		//fmt.Println("house:", house, "current:", current)
		if reflect.DeepEqual(house,current) {
			return true
		}
	}

	return false
}
