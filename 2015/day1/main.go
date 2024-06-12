package main

import (
	"fmt"
	"embed"
)

//go:embed input.txt
var myFileContent embed.FS

func main() {

	fmt.Println("Day 1: Not Quite Lisp")
	fmt.Println("")

	var (
		floor int
	 	nfloor int
	)

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("input length: ", len(input))

	//parse input
	for index, c := range input {
		if string(c) == "(" {
			floor++
			fmt.Println("floor up")
		} else if string(c) == ")" {
			floor--
			fmt.Println("floor down")
		}
		if floor == -1 && nfloor == 0 {
			nfloor = index+1
		}

		fmt.Println("floor: ", floor)
	}

	fmt.Println("")
	fmt.Println("final floor: ", floor)
	fmt.Println("first time negative floor was reached: ", nfloor)
}
