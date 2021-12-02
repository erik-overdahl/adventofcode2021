package day02

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	inputFile := "../../inputs/002-example.txt"
	input, err := utils.ReadlinesStr(inputFile)
	if err != nil {
		panic(fmt.Errorf("Unable to open file '%s': %v", inputFile, err))
	}
	expected := 150
	actual := Part1(input)
	if actual != expected {
		log.Fatalf("Expected %d, got %d for input '%v'", expected, actual, input)
	}
}
