package day03

import (
	"fmt"
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

func Part2(input []string) int {
	numBits := len(input[0])
	nums := make([]int, len(input))
	for i, s := range input {
		N, err := strconv.ParseUint(s, 2, 64)
		if err != nil {
			panic(err)
		}
		n := int(N)
		nums[i] = n
	}

	g := make([]int, len(nums))
	copy(g, nums)
	for i := numBits - 1; i >= 0 && len(g) > 1; i-- {
		mostCommon := mostCommonBitAt(g, i)
		g = filterBitAt(g, mostCommon, i)
	}
	if len(g) != 1 {
		panic(fmt.Errorf("Failed to get single value for G"))
	}

	s := make([]int, len(nums))
	copy(s, nums)
	for i := numBits - 1; i >= 0 && len(s) > 1; i-- {
		leastCommon := flipBits(mostCommonBitAt(s, i), 1)
		s = filterBitAt(s, leastCommon, i)
	}
	if len(s) != 1 {
		panic(fmt.Errorf("Failed to get single value for S"))
	}
	return g[0] * s[0]
}

func filterBitAt(nums []int, val, pos int) []int {
	matches := []int{}
	for _, n := range nums {
		if bitAt(n, pos) == val {
			matches = append(matches, n)
		}
	}
	return matches
}

func mostCommonBitAt(nums []int, pos int) int {
	halfSize := float64(len(nums)) / 2.0
	count := 0
	for _, n := range nums {
		count += bitAt(n, pos)
	}
	if float64(count) >= halfSize {
		return 1
	} else {
		return 0
	}
}

func bitAt(n int, pos int) int {
	return (n >> pos) & 1
}

func flipBits(num int, size int) int {
	mask := (1 << size) - 1
	return (^num) & mask
}
