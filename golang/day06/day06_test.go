package day06

import (
	"testing"
)

const testFile = "../../inputs/006-example.txt"

var testSolution Solution

func TestPopulationOnDay(t *testing.T) {
	if testSolution.input == nil {
		testSolution.Init(testFile)
	}
	cases := [][]int{
		{18, 26},
		{80, 5934},
	}
	for _, c := range cases {
		counts := readInput(testSolution.input)
		actual := populationOnDay(c[0], counts)
		if actual != c[1] {
			t.Fatalf("Expected %d after day %d, got %d", c[1], c[0], actual)
		}
	}
}
