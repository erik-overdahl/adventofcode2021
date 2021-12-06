package day05

import (
	"aoc2021/utils"
	"fmt"
)

type Solution struct {
	input []string
	lines []Line
}

func (d *Solution) Day() int {
	return 5
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
	d.lines = readInput(d.input)
}

func (d *Solution) Part1() string {
	covered := pointsCoveredVerticalHorizontal(d.lines)
	answer := numPointsOverlapped(covered)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	covered := pointsCovered(d.lines)
	answer := numPointsOverlapped(covered)
	return fmt.Sprintf("%d", answer)
}
