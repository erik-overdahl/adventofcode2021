package day02

import (
	"aoc2021/utils"
	"fmt"
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

func (d *Solution) Part1() string {
	if d.position == nil {
		d.position = calculatePosition(d.input)
	}
	answer := d.position.horiz * d.position.aim
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	if d.position == nil {
		d.position = calculatePosition(d.input)
	}
	answer := d.position.horiz * d.position.depth
	return fmt.Sprintf("%d", answer)
}
