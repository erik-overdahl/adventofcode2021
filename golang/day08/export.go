package day08

import (
	"aoc2021/utils"
	"fmt"
)

type Solution struct {
	input []string
}

func (d *Solution) Day() int {
	return 8
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
}

func (d *Solution) Part1() string {
	answer := part1(d.input)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	answer := part2(d.input)
	return fmt.Sprintf("%d", answer)
}
