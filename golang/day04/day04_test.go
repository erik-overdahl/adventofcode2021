package day04

import (
	"aoc2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadlinesStr("../../inputs/004-example.txt")
	if err != nil {
		panic(err)
	}
	expected := 4512
	actual := Part1(input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
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
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{}},
			expected: false,
		},
		{
			name:     "First row marked",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
			expected: true,
		},
		{
			name:     "Some true, no complete row",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{true, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, true, true, false, false}},
			expected: false,
		},
	}
	for _, c := range cases {
		actual := c.card.CheckRows()
		if actual != c.expected {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestCheckCols(t *testing.T) {
	cases := []testCard{
		{
			name:     "All false",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{}},
			expected: false,
		},
		{
			name:     "First col marked",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
			expected: true,
		},
		{
			name:     "Some true, no complete col",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{false, true, true, false, false, true, false, true, false, false, true, true, false, false, false, true, true, true, false, false, true, true, true, false, false}},
			expected: false,
		},
	}
	for _, c := range cases {
		actual := c.card.CheckRows()
		if actual != c.expected {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestSumUnmarked(t *testing.T) {
	cases := []testCard{
		{
			name:     "All false",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{}},
			expected: false,
		},
		{
			name:     "First col marked",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
			expected: true,
		},
		{
			name:     "Some true, no complete col",
			card:     &BingoCard{nums: [25]int8{}, marked: [25]bool{false, true, true, false, false, true, false, true, false, false, true, true, false, false, false, true, true, true, false, false, true, true, true, false, false}},
			expected: false,
		},
	}
	for _, c := range cases {
		actual := c.card.CheckRows()
		if actual != c.expected {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}
