package main

import (
	"embed"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var myFileContent embed.FS

// input.txt content
var content []string

// contains all wires
var wireSlice map[string]uint16

// to sort
var keys []string

// searched key
var searchedKey string

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

	var fields []string
	var binary uint16
	var result uint16
	var isValid bool

	wireSlice = make(map[string]uint16)

	searchedKey = "a"

	for wireSlice[searchedKey] == 0 {
		for _, line := range content {

			fields = strings.Fields(line)

			if len(fields) > 0 {
				if wireSlice[fields[len(fields)-1]] == 0 {
					switch len(fields) {
					case 3:
						result, isValid = getKey(wireSlice, fields[0])

					case 4:
						binary, isValid = getKey(wireSlice, fields[1])
						if isValid == true {
							result = ^binary
						}
					case 5:
						isValid = false
						switch {
						case strings.Contains(line, "OR"):
							result, isValid = bitwise("OR", fields[0], fields[2])
						case strings.Contains(line, "AND"):
							result, isValid = bitwise("AND", fields[0], fields[2])
						case strings.Contains(line, "LSHIFT"):
							result, isValid = bitwise("LSHIFT", fields[0], fields[2])
						case strings.Contains(line, "RSHIFT"):
							result, isValid = bitwise("RSHIFT", fields[0], fields[2])
						}
					}

					if isValid == true {
						wireSlice[fields[len(fields)-1]] = result
					}
				}
			}
		}
	}

	for k := range wireSlice {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if wireSlice[key] != 0 {
			fmt.Println("Key:", key, "Value:", wireSlice[key])
		}
	}
}

func getKey(myMap map[string]uint16, s string) (uint16, bool) {
	var isValid bool
	var result uint16

	key, ok := myMap[s]
	if !ok {
		tmp, err := strconv.ParseUint(s, 10, 16)
		if err == nil {
			isValid = true
			result = uint16(tmp)
		} else {
			isValid = false
		}
	} else {
		isValid = true
		result = key
	}

	return result, isValid
}

func bitwise(operation string, first string, second string) (uint16, bool) {
	var result uint16
	var isValid bool

	var operands []uint16
	params := []string{first, second}

	for _, item := range params {
		result, isValid = getKey(wireSlice, item)
		if isValid == true {
			operands = append(operands, result)
		} else {
			return 0, isValid
		}
	}

	//fmt.Println(first, "value", operands[0])
	//fmt.Println(second, "value", operands[1])

	switch operation {
	case "OR":
		result = operands[0] | operands[1]
	case "AND":
		result = operands[0] & operands[1]
	case "LSHIFT":
		result = operands[0] << operands[1]
	case "RSHIFT":
		result = operands[0] >> operands[1]
	}

	//fmt.Println("operation", operation, "result", result)

	//fmt.Println("wireSlice[fields[0]] | wireSlice[fields[2]]", wireSlice[fields[0]]|wireSlice[fields[2]])
	return result, isValid
}
