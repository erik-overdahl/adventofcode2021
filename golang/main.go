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
	for i, f := range [](func([]int) int){day01.Part1, day01.Part2} {

		solution := f(input)
		fmt.Printf("Day 1 Part %d solution: %d\n", i, solution)
	}
}
