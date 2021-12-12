package day12

import (
	"aoc2021/utils"
)

type Solution struct {
	input []string
}

func (d *Solution) Day() int {
	return 12
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
}

func (d *Solution) Part1() int {
	return part1(d.input)
}

func (d *Solution) Part2() int {
	return 0
}
