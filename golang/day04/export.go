package day04

import (
	"aoc2021/utils"
	"fmt"
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

func (d *Solution) Part1() string {
	draws, cards := readInput(d.input)
	answer := part1(draws, cards)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	draws, cards := readInput(d.input)
	answer := part2(draws, cards)
	return fmt.Sprintf("%d", answer)
}
