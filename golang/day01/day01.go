package day01

func Part1(input []int) int {
	size := len(input)
	numIncreasing := 0
	for i := 0; i < size - 1; i++ {
		if isMonotonicallyIncreasing(input[i:i+2]) {
			numIncreasing++
		}
	}
	return numIncreasing
}

func isMonotonicallyIncreasing(nums []int) bool {
	size := len(nums)
	for i := 0; i < size - 1; i++ {
		if nums[i+1] <= nums[i] {
			return false
		}
	}
	return true
}
