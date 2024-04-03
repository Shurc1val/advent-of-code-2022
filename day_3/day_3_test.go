package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input_string := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	var expected_output int64 = 157

	output := part_one(input_string)
	if output != expected_output {
		t.Log("Expected output", expected_output)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}


func TestPartTwo(t *testing.T) {
	input_string := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	var expected_output int64 = 70
	
	output := part_two(input_string)
	if output != expected_output {
		t.Log("Expected output", expected_output)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}
