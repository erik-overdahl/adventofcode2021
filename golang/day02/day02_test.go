package day02

import (
	"log"
	"testing"
)

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/002-example.txt"
		testSolution.Init(testFile)
	}
	expected := "150"
	actual := testSolution.Part1()
	if actual != expected {
		log.Fatalf("Expected %s, got %s for input '%v'", expected, actual, testSolution.input)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/002-example.txt"
		testSolution.Init(testFile)
	}
	expected := "900"
	actual := testSolution.Part2()
	if actual != expected {
		log.Fatalf("Expected %s, got %s for input '%v'", expected, actual, testSolution.input)
	}
}
