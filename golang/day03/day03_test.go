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
	expected := int(198)
	actual := Part1(input)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
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
