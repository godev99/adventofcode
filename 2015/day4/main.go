package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"encoding/hex"
)

func main() {

	fmt.Println("Day 4: The Ideal Stocking Stuffer")
	fmt.Println("")

	start := "ckczppom"
	hasher := md5.New()
	var first4char, first5char int
	var md5 string

	for i := 0; i < 10000000000000; i++ {

		hasher.Reset()
		hasher.Write([]byte(start + strconv.Itoa(i)))
		md5 = hex.EncodeToString(hasher.Sum(nil))

		if md5[0] == '0' && md5[1] == '0' && md5[2] == '0' && md5[3] == '0' && md5[4] == '0' && first4char == 0 {
			first4char = i
		}

		// check if first character of md5 is equal to 0
		if md5[0] == '0' && md5[1] == '0' && md5[2] == '0' && md5[3] == '0' && md5[4] == '0' && md5[5] == '0'{
			first5char = i
			break
		}
	}

	fmt.Println("First 4 characters: ", first4char)
	fmt.Println("First 5 characters: ", first5char)
}
