package day01

func Part1(input []int) int {
	size := len(input)
	numIncreasing := 0
	for i := 0; i < size - 1; i++ {
		if input[i+1] > input[i] {
			numIncreasing++
		}
	}
	return numIncreasing
}

func Part2(input []int) int {
	numIncreasingWindows := 0
	sum := 0
	for _, n := range input[:3] {
		sum += n
	}
	size := len(input) - 3
	for i := 0; i < size; i++ {
		lastSum := sum
		sum -= input[i]
		sum += input[i+3]
		if sum > lastSum {
			numIncreasingWindows++
		}
	}
	return numIncreasingWindows
}
