package day01

func Part1(input []int) int {
	return numIncreasingWindows(input, 1)
}

func Part2(input []int) int {
	return numIncreasingWindows(input, 3)
}

func numIncreasingWindows(input []int, windowSize int) int {
	size := len(input)
	numIncreasing := 0
	for i := windowSize; i < size; i++ {
		if input[i] > input[i-windowSize] {
			numIncreasing++
		}
	}
	return numIncreasing
}
