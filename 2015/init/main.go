package main

import (
	"embed"
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
		//slog.Error("Error reading file:", err)
		return
	}

	// split content by line
	content = strings.Split(string(input), "\n")
}

func main() {

}
