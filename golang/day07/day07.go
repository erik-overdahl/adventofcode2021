package day07

func minFuelToAlign(crabs map[int]int) int {
	minPos := 100000
	maxPos := 0
	for c, _ := range crabs {
		if c > maxPos {
			maxPos = c
		} else if c < minPos {
			minPos = c
		}
	}
	minFuel := 100000000000
	for i := minPos; i <= maxPos; i++ {
		fuel := fuelToAlign(i, crabs)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func fuelToAlign(position int, crabs map[int]int) int {
	fuel := 0
	for crab, count := range crabs {
		fuel += abs(position-crab) * count
	}
	return fuel
}

func countCrabs(crabs []int) map[int]int {
	counts := make(map[int]int)
	for _, c := range crabs {
		if _, exists := counts[c]; exists {
			counts[c]++
		} else {
			counts[c] = 1
		}
	}
	return counts
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
