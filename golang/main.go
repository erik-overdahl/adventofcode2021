package main

import (
	"aoc2021/day01"
	"aoc2021/utils"
	"fmt"
)

func main() {
	filename := "../inputs/001.txt"
	input, err := utils.ReadlinesInt(filename)
	if err != nil {
		panic(err)
	}
	solution := day01.Part1(input)
	fmt.Printf("Day 1 Part 1 solution: %d\n", solution)
}
