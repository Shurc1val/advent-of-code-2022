package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractRangePairs(inputString string) [][][]int64 {
	rangePairStrings := strings.Split(inputString, "\n")
	rangePairs := [][][]int64{}
	for _, rangePairString := range rangePairStrings {
		ranges := strings.Split(rangePairString, ",")
		rangePair := make([][]int64, 2)
		for i, rString := range ranges {
			r := make([]int64, 2)
			for j, num := range strings.Split(rString, "-") {
				r[j], _ = strconv.ParseInt(num, 10, 64)
			}
			rangePair[i] = r
		}
		rangePairs = append(rangePairs, rangePair)
	}
	return rangePairs
}

func partOne(inputString string) int64 {
	rangePairs := extractRangePairs(inputString)
	var count int64 = 0
	for _, rangePair := range rangePairs {
		if (rangePair[0][0] <= rangePair[1][0] && rangePair[0][1] >= rangePair[1][1]) ||
		(rangePair[0][0] >= rangePair[1][0] && rangePair[0][1] <= rangePair[1][1]) {
			count += 1
		}
	}
	return count
}


func partTwo(inputString string) int {
	rangePairs := extractRangePairs(inputString)
	var count int = len(rangePairs)
	for _, rangePair := range rangePairs {
		if (rangePair[1][0] > rangePair[0][1]) || (rangePair[0][0] > rangePair[1][1]) {
			count -= 1
		}
	}
	return count
}


func main() {
	puzzleInput, _ := os.ReadFile("puzzle_input.txt")
	fmt.Println("Part one:", partOne(string(puzzleInput)))
	fmt.Println("Part two:", partTwo(string(puzzleInput)))
}