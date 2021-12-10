package day10

import (
	"aoc2021/utils"
	"testing"
)

var testFile = "../inputs/010-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		blob := utils.ReadFileToString(testFile)
		testSolution.Init(blob)
	}
	expected := 26397
	actual := part1(testSolution.input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
