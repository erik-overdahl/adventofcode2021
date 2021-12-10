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
	expected := 26
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 61229
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestDecodeOutputDigits(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	allInputs := readInput(testSolution.input)
	allExpected := []int{
		8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315,
	}
	for i, m := range allInputs {
		expected := allExpected[i]
		actual := decode(m)
		if actual != expected {
			t.Errorf("Expected %d, got %d for %s", expected, actual, testSolution.input[i])
		}
	}

}

func TestConvertToByte(t *testing.T) {
	type testcase struct {
		input    string
		expected byte
	}
	cases := []testcase{
		{"a", 0b1},
		{"c", 0b100},
		{"abcdefg", 0b01111111},
	}

	for _, c := range cases {
		actual := convertToByte(c.input)
		if actual != c.expected {
			t.Fatalf("For %s: expected %b, got %b", c.input, c.expected, actual)
		}
	}

}

func TestCountSetBits(t *testing.T) {
	type testcase struct {
		input    byte
		expected int
	}
	cases := []testcase{
		{0b1, 1},
		{0b10, 1},
		{0b11, 2},
		{0b1010, 2},
	}
	for _, c := range cases {
		actual := countSetBits(c.input)
		if actual != c.expected {
			t.Fatalf("For %b: expected %d, got %d", c.input, c.expected, actual)
		}
	}

}
