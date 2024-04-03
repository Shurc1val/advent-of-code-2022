package main

import (
	"fmt"
	"os"
	"strings"
)


func part_one(input_string string) int64 {
	bag_contents := strings.Split(input_string, "\n")
	var score int64
	for _, bag_content := range bag_contents {
		compartment_1 := bag_content[:len(bag_content)/2]
		compartment_2 := bag_content[len(bag_content)/2:]
		for _, item := range compartment_1 {
			if strings.ContainsRune(compartment_2, item) {
				unicode_val := int64(item)
				if unicode_val < 91 {
					score += unicode_val - 38
				} else {
					score += unicode_val - 96
				}
				break
			}
		}
	}
	return score
}


func part_two(input_string string) int64 {
	bag_contents := strings.Split(input_string, "\n")
	var score int64
	for i := 0; i < len(bag_contents) - 1; i += 3 {
		for _, item := range bag_contents[i] {
			if strings.ContainsRune(bag_contents[i+1], item) && strings.ContainsRune(bag_contents[i+2], item) {
				unicode_val := int64(item)
				if unicode_val < 91 {
					score += unicode_val - 38
				} else {
					score += unicode_val - 96
				}
				break
			}
		}
	}
	return score
}


func main() {
	puzzle_input, _ := os.ReadFile("puzzle_input.txt")
	fmt.Println("Part one:", part_one(string(puzzle_input)))
	fmt.Println("Part two:", part_two(string(puzzle_input)))
}