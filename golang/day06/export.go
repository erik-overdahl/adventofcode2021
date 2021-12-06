package day06

import (
	"aoc2021/utils"
	"fmt"
)

type Solution struct {
	input []string
}

func (d *Solution) Day() int {
	return 6
}

func (d *Solution) Init(filename string) {
	d.input = utils.ReadlinesStr(filename)
}

func (d *Solution) Part1() string {
	counts := readInput(d.input)
	answer := populationOnDay(80, counts)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	counts := readInput(d.input)
	answer := populationOnDay(256, counts)
	return fmt.Sprintf("%d", answer)
}
