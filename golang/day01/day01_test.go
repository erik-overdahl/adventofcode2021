package day01

import (
	"aoc2021/utils"
	"testing"
)

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

func TestPart2(t *testing.T) {
	input, err := utils.ReadlinesInt("../../inputs/001-example.txt")
	if err != nil {
		panic(err)
	}
	expected := 5
	actual := Part2(input)
	if actual != expected {
		t.Fatalf("Input %v: expected %d, got %d", input, expected, actual)
	}
}
