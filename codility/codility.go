package main

import (
	"fmt"
	"strconv"
)

func main() {
	a1 := findLongestBinaryGap(333333332)
	fmt.Println(a1)
}

//should use bit shift
func findLongestBinaryGap(number int64) int {
	fmt.Println(strconv.FormatInt(number, 2))
	var maxGapLength, gapLength, count int = 0,0,0

	for _,v:= range strconv.FormatInt(number, 2) {
		if string(v) == "1" {
			gapLength = count
			count = 0
			if gapLength > maxGapLength {
				maxGapLength = gapLength
			}
		} else if string(v) == "0" {
			count++
		}
	}

	return maxGapLength
}