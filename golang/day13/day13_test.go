package day13

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/013-example.txt"

func TestPart1(t *testing.T) {
	testFileBlob := utils.ReadFileToString(testFile)
	lines := utils.ReadlinesStr(testFileBlob)
	points, reflections := readInput(lines)
	points = distinct(applyReflections(points, reflections[:1]))
	expected := 17
	actual := len(points)
	if actual != expected {
		t.Fatalf("Expected %d, got %d:\n%v", expected, actual, points)
	}
}
