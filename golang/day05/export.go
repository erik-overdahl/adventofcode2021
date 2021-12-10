package day05

import (
	"aoc2021/utils"
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

func (d *Solution) Part1() int {
	covered := pointsCoveredVerticalHorizontal(d.lines)
	return numPointsOverlapped(covered)
}

func (d *Solution) Part2() int {
	covered := pointsCovered(d.lines)
	return numPointsOverlapped(covered)
}
