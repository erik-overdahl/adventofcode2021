package day01

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
