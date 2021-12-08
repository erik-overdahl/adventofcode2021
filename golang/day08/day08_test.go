package day08

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/008-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := "26"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}
