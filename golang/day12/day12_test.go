package day12

import (
	"aoc2021/utils"
	"testing"
)

var testFiles = []string{"../inputs/012-exampleA.txt", "../inputs/012-exampleB.txt", "../inputs/012-exampleC.txt"}

var testSolution Solution

func TestPart1(t *testing.T) {
	testGraphs := make([][]string, len(testFiles))
	for i, f := range testFiles {
		testFileBlob := utils.ReadFileToString(f)
		testGraphs[i] = utils.ReadlinesStr(testFileBlob)
	}
	expected := []int{10, 19, 226}
	for i, g := range testGraphs {
		actual := part1(g)
		if actual != expected[i] {
			t.Errorf("Graph %d: Expected %d, got %d", i+1, expected[i], actual)
		}
	}
}
