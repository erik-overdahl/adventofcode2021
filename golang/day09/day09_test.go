package day09

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/009-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 15
	actual := part1(testSolution.input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 1134
	actual := part2(testSolution.input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
