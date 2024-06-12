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
	fmt.Println("Part2")
	fmt.Println("")

	var absX, absY, santaX, santaY, roboSantaX, roboSantaY int
	var isAlreadyVisited bool

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// init matrix with input length
	fmt.Println("input length: ", len(input)/2)
	length := len(input)
	var houses = make([][]int, length+1)

	// first mouvement is a present for first leaving position
	houses[0] = []int{0, 0}
	presents := 1
	fmt.Println("first house visited at:", absX, absY)

	//parse input
	for index, c := range input {
		if index % 2 == 0 {
			absX, absY = move(&santaX, &santaY, string(c))
		} else {
			absX, absY = move(&roboSantaX, &roboSantaY, string(c))
		}

		// check if the house is already visited
		// if not, add it to the matrix
		isAlreadyVisited = alreadyVisited(&houses, []int{absX, absY})
		if !isAlreadyVisited {
			presents++
			fmt.Println("new house visited at:", absX, absY)
		}

		houses[index+1] = []int{absX, absY}
	}

	fmt.Println("")
	fmt.Println("house visited:", presents)
}

// check if the house is already visited
func alreadyVisited(matrix *[][]int, current []int) bool {
	//fmt.Println("matrix:", matrix, "current:", current)
	for _, house := range *matrix {
		//fmt.Println("house:", house, "current:", current)
		if reflect.DeepEqual(house,current) {
			return true
		}
	}

	return false
}

// movement based on direction and index, odd or even
func move(absX *int, absY *int, direction string) (int, int) {
		switch direction {
			case "^":
				*absY++
			case "v":
				*absY--
			case ">":
				*absX++
			case "<":
				*absX--
			default:
		}
		return *absX, *absY
}
