package day07

import (
	"strconv"
	"strings"
)

type Solution struct {
	positions []int
}

func (d *Solution) Day() int {
	return 7
}

func (d *Solution) Init(inputBlob string) {
	sNums := strings.Split(inputBlob, ",")
	d.positions = make([]int, len(sNums))
	for i, s := range sNums {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		d.positions[i] = n
	}
}

func (d *Solution) Part1() int {
	return minFuelToAlign(d.positions, fuelToAlign)
}

func (d *Solution) Part2() int {
	return minFuelToAlign(d.positions, adjustedFuelToAlign)
}
