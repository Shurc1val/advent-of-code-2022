package main

import (
	"fmt"
	"os"
	"strings"
)

func get_result_score(opponent_hand string, self_hand string) int64 {
	hands_string := opponent_hand + " " + self_hand
	switch hands_string {
	case "Rock Paper", "Paper Scissors", "Scissors Rock":
		return 6
	case "Paper Rock", "Scissors Paper", "Rock Scissors":
		return 0
	default:
		return 3
	}
}


func part_one(input_string string) int64 {
	opponent_key := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	self_key := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	shape_scores := map[string]int64{
		"Rock": 1,
		"Paper": 2,
		"Scissors": 3,
	}

	var score int64
	for _, line := range strings.Split(input_string, "\n") {
		encrypted_shapes := strings.Split(line, " ")
		opponent_shape := opponent_key[encrypted_shapes[0]]
		self_shape := self_key[encrypted_shapes[1]]
		score += (shape_scores[self_shape] + get_result_score(opponent_shape, self_shape))
		}
	return score
}

func get_required_move(opponent_hand string, desired_outcome string) string {
	hands_string := opponent_hand + " " + desired_outcome
	switch hands_string {
	case "Rock Draw", "Paper Lose", "Scissors Win":
		return "Rock"
	case "Rock Win", "Paper Draw", "Scissors Lose":
		return "Paper"
	default:
		return "Scissors"
	}
}

func part_two(input_string string) int64 {
	opponent_key := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	result_key := map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}
	result_scores := map[string]int64{
		"Win": 6,
		"Draw": 3,
		"Lose": 0,
	}
	shape_scores := map[string]int64{
		"Rock": 1,
		"Paper": 2,
		"Scissors": 3,
	}

	var score int64
	for _, line := range strings.Split(input_string, "\n") {
		encrypted := strings.Split(line, " ")
		opponent_shape := opponent_key[encrypted[0]]
		result := result_key[encrypted[1]]
		self_shape := get_required_move(opponent_shape, result)
		score += (shape_scores[self_shape] + result_scores[result])
		}
	return score
}


func main() {
	puzzle_input, _ := os.ReadFile("puzzle_input.txt")
	fmt.Println(part_one(string(puzzle_input)))
	fmt.Println(part_two(string(puzzle_input)))
}