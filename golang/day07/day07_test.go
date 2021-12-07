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

func TestPart2(t *testing.T) {
	if testSolution.positions == nil {
		testBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testBlob)
	}
	expected := "168"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestAdjustedFuelToAlign(t *testing.T) {
	if testSolution.positions == nil {
		testBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testBlob)
	}
	testCrabs := countCrabs(testSolution.positions)
	cases := [][]int{
		{2, 206},
		{5, 168},
	}
	for _, c := range cases {
		actual := adjustedFuelToAlign(c[0], testCrabs)
		if actual != c[1] {
			t.Fatalf("Position %d: Expected %d, got %d", c[0], c[1], actual)
		}
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

func TestSumTo(t *testing.T) {
	cases := [][]int{
		{10, 55},
		{11, 66},
	}
	for _, c := range cases {
		actual := sumTo(c[0])
		if actual != c[1] {
			t.Fatalf("n= %d: Expected %d, got %d", c[0], c[1], actual)
		}
	}
}
