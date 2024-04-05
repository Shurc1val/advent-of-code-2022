package main

import "testing"

func TestPartOne(t *testing.T) {
	inputString := ``
	var expectedOutput int64 = 0

	output := partOne(inputString)
	if output != expectedOutput {
		t.Log("Expected output", expectedOutput)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}

/*
func TestPartTwo(t *testing.T) {
	inputString := ``
	var expectedOutput int64 = 0
	
	output := partTwo(inputString)
	if output != expectedOutput {
		t.Log("Expected output", expectedOutput)
		t.Log("Actual output", output)
		t.Fatalf("Incorrect")
	}
}
*/