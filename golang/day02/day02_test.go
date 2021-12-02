package day02

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"testing"
)

func getData() []string {
	inputFile := "../../inputs/002-example.txt"
	input, err := utils.ReadlinesStr(inputFile)
	if err != nil {
		panic(fmt.Errorf("Unable to open file '%s': %v", inputFile, err))
	}
	return input
}

func TestPart1(t *testing.T) {
	input := getData()
	expected := 150
	actual := Part1(input)
	if actual != expected {
		log.Fatalf("Expected %d, got %d for input '%v'", expected, actual, input)
	}
}

func TestPart2(t *testing.T) {
	input := getData()
	expected := 900
	actual := Part2(input)
	if actual != expected {
		log.Fatalf("Expected %d, got %d for input '%v'", expected, actual, input)
	}
}
