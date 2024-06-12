package main

import (
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var myFileContent embed.FS

func main() {

	fmt.Println("Day 5: Doesn't He Have Intern-Elves For This?")
	fmt.Println("Part2")
	fmt.Println("")

	var isContainsPairs bool
	var niceStrings int
	var pairs string

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// split content by line
	content := strings.Split(string(input), "\n")

	for _, line := range content {
		// check if it contains a pair of any two letters that appears at least twice in the string without overlapping
		isContainsPairs, pairs = checkPairs(line)

		// check if it contains at least one letter which repeats with exactly one letter between them
		if isContainsPairs {
			if unique, nice := checkUniqueChar(line); unique {
				fmt.Println("line:", line, "pairs:", pairs, "unique:", nice, "=> nice string")
				niceStrings++
			} else {
				fmt.Println("line:", line, "pairs:", pairs, "=> bad string")
			}
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("Nice strings:", niceStrings)
}

func checkPairs(line string) (bool, string) {

	var r *regexp.Regexp
	var indexes [][]int
	var regex string

	for i := 0; i < len(line)-1; i++ {
		regex = string(line[i]) + string(line[i+1])
		r, _ = regexp.Compile(regex)
		indexes = r.FindAllStringSubmatchIndex(line, 2)

		if len(indexes) > 1 {
			return true, string(line[indexes[0][0]:indexes[0][1]])
		}
	}

	return false, ""
}

func checkUniqueChar(remaining string) (bool, string) {
	var left, right, current string
	for in := 1; in <= len(remaining)-2; in++ {

		left = string(remaining[in-1])
		right = string(remaining[in+1])
		current = string(remaining[in])

		if left == right {
			return true, left + current + right
		}

	}
	return false, ""
}
