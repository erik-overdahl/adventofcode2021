package day03

import (
	"testing"
)

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/003-example.txt"
		testSolution.Init(testFile)
	}
	expected := "198"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testFile := "../../inputs/003-example.txt"
		testSolution.Init(testFile)
	}
	expected := "230"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}
func TestMostCommonBitAt(t *testing.T) {
	type testcase struct {
		input    []int
		pos      int
		expected int
	}
	cases := []testcase{
		{input: []int{0b10, 0b10, 0b10, 0b01, 0b01}, pos: 1, expected: 1},
		{input: []int{0b10, 0b10, 0b01, 0b01}, pos: 1, expected: 1},
		{input: []int{0b10, 0b10, 0b10, 0b01, 0b01}, pos: 0, expected: 0},
	}
	for _, c := range cases {
		actual := mostCommonBitAt(c.input, c.pos)
		if actual != c.expected {
			t.Fatalf("expected %d, got %d for input %v at pos %d", c.expected, actual, c.input, c.pos)
		}
	}
}

func TestBitAt(t *testing.T) {
	type testcase struct {
		input    int
		pos      int
		expected int
	}
	cases := []testcase{
		{int(0b100), 0, 0},
		{int(0b101), 0, 1},
		{int(0b101), 2, 1},
		{int(0b101), 3, 0},
	}
	for _, c := range cases {
		actual := bitAt(c.input, c.pos)
		if actual != c.expected {
			t.Fatalf("%b at %d: expected %d, got %d", c.input, c.pos, c.expected, actual)
		}
	}
}
