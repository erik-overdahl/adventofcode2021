package day01

import (
	"aoc2021/utils"
)

type Solution struct {
	input []int
}

func (d *Solution) Day() int {
	return 1
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesInt(inputBlob)
}

func (d *Solution) Part1() int {
	return numIncreasingWindows(d.input, 1)
}

func (d *Solution) Part2() int {
	return numIncreasingWindows(d.input, 3)
}
