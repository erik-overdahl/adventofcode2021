package day11

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/011-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		blob := utils.ReadFileToString(testFile)
		testSolution.Init(blob)
	}
	cases := [][]int{
		{10, 204},
		{100, 1656},
	}
	for _, c := range cases {
		expected := c[1]
		actual := part1(testSolution.input, c[0])
		if actual != expected {
			t.Fatalf("After %d generations: expected %d, got %d", c[0], expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		blob := utils.ReadFileToString(testFile)
		testSolution.Init(blob)
	}
	expected := 195
	actual := part2(testSolution.input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
