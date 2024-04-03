package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func get_elf_loads(calorie_string string) []int64 {
	elf_inventories := strings.Split(calorie_string, "\n\n")
	var elf_loads []int64
	var items []string
	var num int64
	var sum int64
	for _, inventory := range elf_inventories {
		items = strings.Split(inventory, "\n")
		sum = 0
		for _, item := range items {
			num, _ = strconv.ParseInt(item, 10, 0)
			sum += num
		}
		elf_loads = append(elf_loads, sum)
	}

	slices.Sort(elf_loads)
	return elf_loads
}


func part_one(calorie_string string) int64 {
	return slices.Max(get_elf_loads(calorie_string))
}


func part_two(calorie_string string) int64 {
	elf_loads := get_elf_loads(calorie_string)
	slices.Sort(elf_loads)
	var val int64
	for i := range 3 {
		val += elf_loads[len(elf_loads) - i - 1]
	}
	return val
}


func main() {
	puzzle_input, _ := os.ReadFile("puzzle_input.txt")
	fmt.Println(part_one(string(puzzle_input)))
	fmt.Println(part_two(string(puzzle_input)))
}