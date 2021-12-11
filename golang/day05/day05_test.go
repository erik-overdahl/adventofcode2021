package day05

import (
	"aoc2021/utils"
	"fmt"
	"testing"
)

const testFile = "../inputs/005-example.txt"

var testSolution Solution

func TestPart1(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	expected := 5
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
	expected := 12
	actual := testSolution.Part2()
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func parseExpectedOverlaps(reference []string) []byte {
	expectedMap := make([]byte, len(reference)*len(reference[0]))
	for y, s := range reference {
		for x, c := range s {
			idx := (y * len(reference[0])) + x
			if c != '.' {
				expectedMap[idx] = byte(c) - 0b00110000
			} else {
				expectedMap[idx] = 0
			}
		}
	}
	return expectedMap
}

func makeVisual(reference []string, actual []byte) []string {
	actualVisual := make([]string, len(reference))
	for y, row := range reference {
		newRow := make([]byte, len(row))
		for x := range row {
			n := actual[(y*len(reference[0]))+x]
			if n == 0 {
				newRow[x] = '.'
			} else {
				newRow[x] = n + 0b00110000
			}
		}
		actualVisual[y] = string(newRow)
	}
	return actualVisual
}

func visualToString(visual []string) string {
	var out string
	for _, s := range visual {
		out = fmt.Sprintf("%s%s\n", out, s)
	}
	return out
}

func TestPointsCoveredVerticalHorizontal(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	comparison := []string{
		".......1..",
		"..1....1..",
		"..1....1..",
		".......1..",
		".112111211",
		"..........",
		"..........",
		"..........",
		"..........",
		"222111....",
	}
	expected := parseExpectedOverlaps(comparison)
	actual := pointsCoveredVerticalHorizontal(testSolution.lines)
	visual := makeVisual(comparison, actual)
	for i, n := range expected {
		actualN := actual[i]
		if actualN != n {
			y := i / len(comparison[0])
			x := i % len(comparison[0])
			msg := fmt.Sprintf("Expected %d, got %d for point %d, %d\n\nexpected:\tactual:\n", n, actualN, x, y)
			for i, row := range comparison {
				msg = fmt.Sprintf("%s%s\t%s\n", msg, row, visual[i])
			}
			t.Fatal(msg)
		}
	}
}

func TestPointsCoveredAll(t *testing.T) {
	if testSolution.input == nil {
		testFileBlob := utils.ReadFileToString(testFile)
		testSolution.Init(testFileBlob)
	}
	comparison := []string{
		"1.1....11.",
		".111...2..",
		"..2.1.111.",
		"...1.2.2..",
		".112313211",
		"...1.2....",
		"..1...1...",
		".1.....1..",
		"1.......1.",
		"222111....",
	}
	expected := parseExpectedOverlaps(comparison)
	actual := pointsCovered(testSolution.lines)
	visual := makeVisual(comparison, actual)
	for i, n := range expected {
		actualN := actual[i]
		if actualN != n {
			y := i / len(comparison[0])
			x := i % len(comparison[0])
			msg := fmt.Sprintf("Expected %d, got %d for point %d, %d\n\nexpected:\tactual:\n", n, actualN, x, y)
			for i, row := range comparison {
				msg = fmt.Sprintf("%s%s\t%s\n", msg, row, visual[i])
			}
			t.Fatal(msg)
		}
	}
}

func TestMakeLine(t *testing.T) {
	type testcase struct {
		p1       Point
		p2       Point
		expected Line
	}
	cases := []testcase{
		{p1: Point{1, 1}, p2: Point{1, 3}, expected: Line{Point{1, 3}, Point{1, 1}}}, //horizontal
		{p1: Point{1, 6}, p2: Point{1, 3}, expected: Line{Point{1, 6}, Point{1, 3}}}, //horizontal
		{p1: Point{3, 3}, p2: Point{6, 3}, expected: Line{Point{3, 3}, Point{6, 3}}}, //vertical
		{p1: Point{4, 0}, p2: Point{0, 0}, expected: Line{Point{0, 0}, Point{4, 0}}}, //vertical
		{p1: Point{0, 0}, p2: Point{4, 4}, expected: Line{Point{0, 0}, Point{4, 4}}}, //positive
		{p1: Point{4, 4}, p2: Point{0, 0}, expected: Line{Point{0, 0}, Point{4, 4}}}, //positive
		{p1: Point{0, 4}, p2: Point{4, 0}, expected: Line{Point{0, 4}, Point{4, 0}}}, //negative
		{p1: Point{4, 0}, p2: Point{0, 4}, expected: Line{Point{0, 4}, Point{4, 0}}}, //negative
	}
	for _, c := range cases {
		actual := makeLine(c.p1, c.p2)
		if actual != c.expected {
			t.Fatalf("Expected %v, got %v for p1 = %v and p2 = %v", c.expected, actual, c.p1, c.p2)
		}
	}
}

func TestReadInput(t *testing.T) {
	input := []string{"1,2 -> 1,1", "5,1 -> 2,1"}
	expected := []Line{
		{Point{1, 2}, Point{1, 1}},
		{Point{2, 1}, Point{5, 1}},
	}
	actual := readInput(input)
	for i, l := range actual {
		if l != expected[i] {
			t.Fatalf("Expected %v, got %v", expected[i], l)
		}
	}
}
