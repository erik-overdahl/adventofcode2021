package day06

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	counts := readInput(input)
	return populationOnDay(80, counts)
}

func Part2(input []string) int {
	counts := readInput(input)
	return populationOnDay(256, counts)
}

func populationOnDay(day int, counts [9]int) int {
	for i := 0; i < day; i++ {
		counts = runGeneration(counts)
	}
	return population(counts)
}

func runGeneration(counts [9]int) [9]int {
	size := len(counts)
	spawning := counts[0]
	for i := 1; i < size; i++ {
		counts[i-1] = counts[i]
	}
	counts[6] += spawning
	counts[8] = spawning
	return counts
}

func population(counts [9]int) int {
	pop := 0
	for _, x := range counts {
		pop += x
	}
	return pop
}

func readInput(input []string) (counts [9]int) {
	s := strings.Split(input[0], ",")
	for _, fish := range s {
		fishTime, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}
		counts[fishTime]++
	}
	return
}
