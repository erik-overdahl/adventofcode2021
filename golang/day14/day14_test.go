package day14

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/014-example.txt"

func TestPart1(t *testing.T) {
	testFileBlob := utils.ReadFileToString(testFile)
	input := utils.ReadlinesStr(testFileBlob)
	expected := 1588
	actual := part1(input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	testFileBlob := utils.ReadFileToString(testFile)
	input := utils.ReadlinesStr(testFileBlob)
	expected := 2188189693529
	actual := part2(input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
