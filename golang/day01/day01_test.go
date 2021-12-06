package day01

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/001-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := "7"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Input %v: expected %s, got %s", testSolution.input, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := "5"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Input %v: expected %s, got %s", testSolution.input, expected, actual)
	}
}
