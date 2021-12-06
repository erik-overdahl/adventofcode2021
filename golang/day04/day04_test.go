package day04

import (
	"testing"
)

const testFile = "../../inputs/004-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testSolution.Init(testFile)
	}
	expected := "4512"
	actual := testSolution.Part1()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	if testSolution.input == nil {
		testSolution.Init(testFile)
	}
	expected := "1924"
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

type testCard struct {
	name     string
	card     *BingoCard
	expected bool
}

func TestCheckRows(t *testing.T) {
	type testcase struct {
		name     string
		card     *BingoCard
		expected bool
	}
	cases := []testcase{
		{
			name:     "All false",
			card:     &BingoCard{nums: [25]int{}, marked: [25]bool{}},
			expected: false,
		},
		{
			name:     "First row marked",
			card:     &BingoCard{nums: [25]int{}, marked: [25]bool{true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
			expected: true,
		},
		{
			name:     "Some true, no complete row",
			card:     &BingoCard{nums: [25]int{}, marked: [25]bool{true, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, true, true, false, false}},
			expected: false,
		},
	}
	for _, c := range cases {
		actual := CheckRows(c.card)
		if actual != c.expected {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestCheckCols(t *testing.T) {
	cases := []testCard{
		{
			name:     "All false",
			card:     &BingoCard{nums: [25]int{}, marked: [25]bool{}},
			expected: false,
		},
		{
			name: "First col marked",
			card: &BingoCard{nums: [25]int{}, marked: [25]bool{
				true, true, true, true, true,
				true, false, false, false, false,
				true, false, false, false, false,
				true, false, false, false, false,
				true, false, false, false, false}},
			expected: true,
		},
		{
			name:     "Some true, no complete col",
			card:     &BingoCard{nums: [25]int{}, marked: [25]bool{false, true, true, false, false, true, false, true, false, false, true, true, false, false, false, true, true, true, false, false, true, true, true, false, false}},
			expected: false,
		},
	}
	for _, c := range cases {
		actual := CheckCols(c.card)
		if actual != c.expected {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}
