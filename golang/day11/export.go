package day11

import "aoc2021/utils"

type Solution struct {
	input [][]byte
}

func (d *Solution) Day() int {
	return 11
}

func (d *Solution) Init(inputBlob string) {
	input := utils.ReadlinesStr(inputBlob)
	lines := make([][]byte, len(input))
	for i, s := range input {
		lines[i] = make([]byte, len(s))
		for j, c := range s {
			lines[i][j] = byte(c) - 0b00110000
		}
	}
	d.input = lines
}

func (d *Solution) Part1() int {
	return part1(d.input, 100)
}

func (d *Solution) Part2() int {
	return part2(d.input)
}
