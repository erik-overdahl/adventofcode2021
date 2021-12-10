package day06

import (
	"aoc2021/utils"
)

type Solution struct {
	input []string
}

func (d *Solution) Day() int {
	return 6
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
}

func (d *Solution) Part1() int {
	counts := readInput(d.input)
	return populationOnDay(80, counts)
}

func (d *Solution) Part2() int {
	counts := readInput(d.input)
	return populationOnDay(256, counts)
}
