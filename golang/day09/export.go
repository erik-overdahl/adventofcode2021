package day09

import (
	"aoc2021/utils"
)

type Solution struct {
	input [][]int
}

func (d *Solution) Day() int {
	return 9
}

func (d *Solution) Init(inputBlob string) {
	lines := utils.ReadlinesStr(inputBlob)
	nums := make([][]int, len(lines))
	for i, l := range lines {
		row := make([]int, len(l))
		for j := range l {
			row[j] = int(byte(l[j]) - 0b0110000)
		}
		nums[i] = row
	}
	d.input = nums
}

func (d *Solution) Part1() int {
	return part1(d.input)
}

func (d *Solution) Part2() int {
	return part2(d.input)
}
