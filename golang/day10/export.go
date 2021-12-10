package day10

import (
	"aoc2021/utils"
)

type Solution struct {
	input [][]byte
}

func (d *Solution) Day() int {
	return 10
}

func (d *Solution) Init(inputBlob string) {
	input := utils.ReadlinesStr(inputBlob)
	lines := make([][]byte, len(input))
	for i, s := range input {
		lines[i] = []byte(s)
	}
	d.input = lines
}

func (d *Solution) Part1() int {
	return part1(d.input)
}

func (d *Solution) Part2() int {
	return part2(d.input)
}
