package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
	positions []int
	crabs     map[int]int
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
	d.crabs = countCrabs(d.positions)
}

func (d *Solution) Part1() string {
	answer := minFuelToAlign(d.crabs, fuelToAlign)
	return fmt.Sprintf("%d", answer)
}

func (d *Solution) Part2() string {
	answer := minFuelToAlign(d.crabs, adjustedFuelToAlign)
	return fmt.Sprintf("%d", answer)
}
