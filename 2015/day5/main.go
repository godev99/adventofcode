package main

import (
	"fmt"
	"embed"
	"strings"
)

//go:embed input.txt
var myFileContent embed.FS

func main() {

	fmt.Println("Day 5: Doesn't He Have Intern-Elves For This?")
	fmt.Println("")

	var contains3vowels, containsTwice, containsForbiddenStrings bool
	var niceStrings int

	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// split content by line
	content := strings.Split(string(input), "\n")

	for _, line := range content {
		// check if it contains at least three vowels (aeiou only)
		// fmt.Println("check if it contains at least three vowels:", line, checkVowels(line))
		contains3vowels = checkVowels(line)

		// check if it contains at least one letter that appears twice in a row
		// fmt.Println("check if it contains at least one letter that appears twice in a row:", line, checkTwice(line))
		containsTwice = checkTwice(line)

		// check if it does not contain the strings ab, cd, pq, or xy
		// fmt.Println("check if it does not contain the strings ab, cd, pq, or xy:", line, checkForbiddenStrings(content))
		containsForbiddenStrings = checkForbiddenStrings(line)

		if contains3vowels && containsTwice && containsForbiddenStrings {
			niceStrings++
		}
	}

	fmt.Println("Nice strings:", niceStrings)
}

func checkVowels(line string) bool {
	count := 0
	for _, letter := range line {
		if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
			count += 1
		}
	}

	if count >= 3 {
		return true
	}

	return false
}

func checkTwice(line string) bool {
	if len(line) > 0 {
		start := line[0]
		for i := 1; i < len(line); i++ {
			if start == line[i] {
				return true
			} else {
				start = line[i]
			}
		}
	}

	return false
}

func checkForbiddenStrings(line string) bool {
	if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		return false
	}
	return true
}
