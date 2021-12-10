package day04

import (
	"aoc2021/utils"
)

type Solution struct {
	input []string
}

func (d *Solution) Day() int {
	return 4
}

func (d *Solution) Init(inputBlob string) {
	d.input = utils.ReadlinesStr(inputBlob)
}

func (d *Solution) Part1() int {
	draws, cards := readInput(d.input)
	return part1(draws, cards)
}

func (d *Solution) Part2() int {
	draws, cards := readInput(d.input)
	return part2(draws, cards)
}
