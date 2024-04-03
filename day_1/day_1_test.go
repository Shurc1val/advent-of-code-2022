package main

import "testing"

func TestPartOne(t *testing.T) {
	input_string := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	var expected_output int64 = 24000
	output := part_one(input_string)
	if output != expected_output {
		t.Fatalf("Incorrect")
	}
}


func TestPartTwo(t *testing.T) {
	input_string := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	var expected_output int64 = 45000
	output := part_two(input_string)
	if output != expected_output {
		t.Fatalf("Incorrect")
	}
}