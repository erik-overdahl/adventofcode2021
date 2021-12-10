package day02

import (
	"aoc2021/utils"
)

type Solution struct {
	input    []string
	position *SubmarinePosition
}

func (d *Solution) Day() int {
	return 2
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
}

func (d *Solution) Part1() int {
	if d.position == nil {
		d.position = calculatePosition(d.input)
	}
	return d.position.horiz * d.position.aim
}

func (d *Solution) Part2() int {
	if d.position == nil {
		d.position = calculatePosition(d.input)
	}
	return d.position.horiz * d.position.depth
}
