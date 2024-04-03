package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input_string := `A Y
B X
C Z`
	var expected_output int64 = 15

	output := part_one(input_string)
	if output != expected_output {
		t.Fatalf("Incorrect")
	}
}


func TestPartTwo(t *testing.T) {
	input_string := `A Y
B X
C Z`
	var expected_output int64 = 12
	output := part_two(input_string)
	if output != expected_output {
		t.Fatalf("Incorrect")
	}
}
