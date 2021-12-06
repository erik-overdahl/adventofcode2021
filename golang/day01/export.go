package day01

import (
	"aoc2021/utils"
	"fmt"
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

func (d *Solution) Part1() string {
	answer := numIncreasingWindows(d.input, 1)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	answer := numIncreasingWindows(d.input, 3)
	return fmt.Sprintf("%d", answer)
}
