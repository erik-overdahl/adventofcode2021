package day03

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := uint64(198)
	actual := Part1(input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
