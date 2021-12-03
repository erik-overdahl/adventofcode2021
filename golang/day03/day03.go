package day03

import (
	"strconv"
)

func Part1(input []string) uint64 {
	halfSize := uint64(len(input) / 2)
	numBits := len(input[0])
	gammaVals := make([]uint64, numBits)
	for _, s := range input {
		n, err := strconv.ParseUint(s, 2, 64)
		if err != nil {
			panic(err)
		}
		for i, _ := range gammaVals {
			gammaVals[i] += (n >> i) & 1
		}
	}
	gamma := uint64(0)
	var bit uint64
	for j := numBits - 1; j >= 0; j-- {
		if gammaVals[j] > halfSize {
			bit = 1
		} else {
			bit = 0
		}
		gamma = (gamma << 1) | bit
	}
	mask := uint64((1 << numBits) - 1)
	epsilon := (^gamma) & mask
	return epsilon * gamma
}

