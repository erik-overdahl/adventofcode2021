package day07

func minFuelToAlign(crabs []int, fuelCalc func(int, []int) int) int {
	minPos := 100000
	maxPos := 0
	for c := range crabs {
		if c > maxPos {
			maxPos = c
		} else if c < minPos {
			minPos = c
		}
	}
	minFuel := (1 << 62)
	for i := minPos; i <= maxPos; i++ {
		fuel := fuelCalc(i, crabs)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func adjustedFuelToAlign(position int, crabs []int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += sumTo(abs(position - crab))
	}
	return fuel
}

func fuelToAlign(position int, crabs []int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += abs(position - crab)
	}
	return fuel
}

func sumTo(n int) int {
	return (n * (n + 1)) >> 1
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
