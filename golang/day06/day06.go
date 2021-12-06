package day06

import (
	"strconv"
	"strings"
)

func populationOnDay(day int, counts []int) int {
	for i := 0; i < day; i++ {
		counts = runGeneration(counts)
	}
	return population(counts)
}

func runGeneration(counts []int) []int {
	size := len(counts)
	spawning := counts[0]
	for i := 1; i < size; i++ {
		counts[i-1] = counts[i]
	}
	counts[6] += spawning
	counts[8] = spawning
	return counts
}

func population(counts []int) int {
	pop := 0
	for _, x := range counts {
		pop += x
	}
	return pop
}

func readInput(input []string) []int {
	s := strings.Split(input[0], ",")
	counts := make([]int, 9)
	for _, fish := range s {
		fishTime, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}
		counts[fishTime]++
	}
	return counts
}
