package day07

import (
	"aoc2021/utils"
	"testing"
)

const testFile = "../inputs/007-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.positions == nil {
		testBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testBlob)
	}
	expected := "37"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestFuelToAlign(t *testing.T) {
	if testSolution.positions == nil {
		testBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testBlob)
	}
	testCrabs := countCrabs(testSolution.positions)
	cases := [][]int{
		{2, 37},
		{1, 41},
		{3, 39},
		{10, 71},
	}
	for _, c := range cases {
		actual := fuelToAlign(c[0], testCrabs)
		if actual != c[1] {
			t.Fatalf("Position %d: Expected %d, got %d", c[0], c[1], actual)
		}
	}
}
