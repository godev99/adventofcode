package main

import (
	"embed"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

//go:embed input.txt
var myFileContent embed.FS

// input.txt content
var content []string

func init() {
	// get input from file
	input, err := myFileContent.ReadFile("input.txt")
	if err != nil {
		slog.Error("Error reading file:", err)
		return
	}

	// split content by line
	content = strings.Split(string(input), "\n")
}

func main() {
	//var realChars int
	//var chars int
	//var result int
	var result []string

	for _, i := range content {
		//chars = getChars(i)
		//realChars = getRealChars(i)
		//result = result + (chars - realChars)
		if i != "" {
			r := regexp.MustCompile(`\\`)
			result = r.FindAllString(i, -1)
			fmt.Println(len(result))
			r = regexp.MustCompile(`\\"`)
			result = r.FindAllString(i, -1)
			fmt.Println(len(result))
			r = regexp.MustCompile("\\[x][a-z0-9]{2}")
			result = r.FindAllString(i, -1)
			fmt.Println(len(result))
		}
	}
}

func getChars(s string) int {
	return len(s)
}

func getRealChars(s string) int {
	r := regexp.MustCompile(`\\`)
	fmt.Println("regex", r.FindAllString(s, -1))

	/*
		for _, i := range s {
			fmt.Println(i)

			//if string(i) == "\\" {
			//	fmt.Println("\\ detected")
			//}
		}
	*/
	return len(s)
}
