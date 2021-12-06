package day01

import (
	"testing"
)

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/001-example.txt"
		testSolution.Init(testFile)
	}
	expected := "7"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Input %v: expected %s, got %s", testSolution.input, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/001-example.txt"
		testSolution.Init(testFile)
	}
	expected := "5"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Input %v: expected %s, got %s", testSolution.input, expected, actual)
	}
}
