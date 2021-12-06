package day06

import (
	"aoc2021/utils"
	"testing"
)

func TestPopulationOnDay(t *testing.T) {
	input, err := utils.ReadlinesStr("../../inputs/006-example.txt")
	if err != nil {
		panic(err)
	}
	counts := readInput(input)
	cases := [][]int{
		{18, 26},
		{80, 5934},
	}
	for _, c := range cases {
		actual := populationOnDay(c[0], counts)
		if actual != c[1] {
			t.Fatalf("Expected %d after day %d, got %d", c[1], c[0], actual)
		}
	}
}
