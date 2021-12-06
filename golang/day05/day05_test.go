package day05

import (
	"testing"
)

const testFile = "../../inputs/005-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testSolution.Init(testFile)
	}
	expected := "5"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testSolution.Init(testFile)
	}
	expected := "12"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestReadInput(t *testing.T) {
	input := []string{"1,2 -> 1,1", "5,1 -> 2,1"}
	expected := []Line{
		{Point{1, 2}, Point{1, 1}},
		{Point{2, 1}, Point{5, 1}},
	}
	actual := readInput(input)
	for i, l := range actual {
		if l != expected[i] {
			t.Fatalf("Expected %v, got %v", expected[i], l)
		}
	}
}
