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

		}
	}
}
