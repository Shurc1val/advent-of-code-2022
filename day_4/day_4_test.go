package main

import "testing"

func TestPartOne(t *testing.T) {
	inputString := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	var expectedOutput int64 = 2

	output := partOne(inputString)
	if output != expectedOutput {
		t.Log("Expected output", expectedOutput)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}

func TestPartTwo(t *testing.T) {
	inputString := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	var expectedOutput int = 4
	
	output := partTwo(inputString)
	if output != expectedOutput {
		t.Log("Expected output", expectedOutput)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}
