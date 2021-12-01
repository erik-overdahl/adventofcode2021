package day01

import (
	"aoc2021/utils"
	"testing"
)

func TestIsMonotonicallyIncreasing(t *testing.T) {
	type testcase struct {
		input	 []int
		expected bool
	}
	largerSlice := []int{1, 2, 3, 2, 1}
	cases := []testcase{
		{input: []int{1, 2, 3}, expected: true},
		{input: []int{2, 2, 3}, expected: false},
		{input: []int{3, 2, 3}, expected: false},
		{input: largerSlice[:3], expected: true},
		{input: largerSlice[1:5], expected: false},
	}
	for _, x := range cases {
		actual := isMonotonicallyIncreasing(x.input)
		if actual != x.expected {
			t.Fatalf("Input %v: expected %v, got %v", x.input, x.expected, actual)
		}
	}
}

func TestPart1(t *testing.T) {
	input, err := utils.ReadlinesInt("../../inputs/001-example.txt")
	if err != nil {
		panic(err)
	}
	expected := 7
	actual := Part1(input)
	if actual != expected {
		t.Fatalf("Input %v: expected %d, got %d", input, expected, actual)
	}
}
