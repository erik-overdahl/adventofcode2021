package day02

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/002-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 150
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %d, got %d for input '%v'", expected, actual, testSolution.input)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 900
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %d, got %d for input '%v'", expected, actual, testSolution.input)
	}
}
