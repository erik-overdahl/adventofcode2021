package day03

import (
	"strconv"
)

func Part1(input []string) int {
	halfSize := len(input) / 2
	numBits := len(input[0])
	gammaVals := make([]int, numBits)
	for _, s := range input {
		N, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			panic(err)
		}
		n := int(N)
		for i := range gammaVals {
			gammaVals[i] += bitAt(n, i)
		}
	}
	gamma := 0
	var bit int
	for j := numBits - 1; j >= 0; j-- {
		if gammaVals[j] > halfSize {
			bit = 1
		} else {
			bit = 0
		}
		gamma = (gamma << 1) | bit
	}
	epsilon := flipBits(gamma, numBits)
	return epsilon * gamma
}

func bitAt(n int, pos int) int {
	return (n >> pos) & 1
}

func flipBits(num int, size int) int {
	mask := (1 << size) - 1
	return (^num) & mask
}
