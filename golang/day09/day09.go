package day09

type point struct {
	x int
	y int
}

func part1(heightmap [][]int) int {
	localMinima := findLocalMinima(heightmap)
	totalRisk := 0
	for _, pt := range localMinima {
		totalRisk += heightmap[pt.y][pt.x] + 1
	}
	return totalRisk
}

func part2(h [][]int) int {
	yMax, xMax := len(h), len(h[0])
	basins := findLocalMinima(h)
	sizes := make([]int, len(basins))
	for i, b := range basins {
		sizes[i] = 1
		val := 100 + i
		h[b.y][b.x] = val
		for points := []point{b}; len(points) > 0; {
			p := points[0]
			points = points[1:]
			neighbors := make([]point, 0, 4)
			switch p.x {
			case 0:
				neighbors = append(neighbors, point{p.x + 1, p.y})
			case xMax - 1:
				neighbors = append(neighbors, point{p.x - 1, p.y})
			default:
				neighbors = append(neighbors, point{p.x + 1, p.y}, point{p.x - 1, p.y})
			}
			switch p.y {
			case 0:
				neighbors = append(neighbors, point{p.x, p.y + 1})
			case yMax - 1:
				neighbors = append(neighbors, point{p.x, p.y - 1})
			default:
				neighbors = append(neighbors, point{p.x, p.y + 1}, point{p.x, p.y - 1})
			}
			for _, n := range neighbors {
				height := h[n.y][n.x]
				if height != 9 && height != val {
					h[n.y][n.x] = val
					points = append(points, n)
					sizes[i]++
				}
			}
		}
	}
	maxes := make([]int, 3)
	for _, size := range sizes {
		if size > maxes[0] {
			maxes[2] = maxes[1]
			maxes[1] = maxes[0]
			maxes[0] = size
		} else if size > maxes[1] {
			maxes[2] = maxes[1]
			maxes[1] = size
		} else if size > maxes[2] {
			maxes[2] = size
		}
	}
	product := 1
	for _, m := range maxes {
		product *= m
	}
	return product
}

func findLocalMinima(heightmap [][]int) []point {
	localMinima := []point{}
	xSize, ySize := len(heightmap[0]), len(heightmap)
	for y, row := range heightmap {
		var localYMin func(x, height int) bool
		switch y {
		case 0:
			nextRow := heightmap[y+1]
			localYMin = func(x, height int) bool { return nextRow[x] > height }
		case ySize - 1:
			prevRow := heightmap[y-1]
			localYMin = func(x, height int) bool { return prevRow[x] > height }
		default:
			prevRow := heightmap[y-1]
			nextRow := heightmap[y+1]
			localYMin = func(x, height int) bool { return nextRow[x] > height && prevRow[x] > height }
		}
		prevHeight := 0
		for x, height := range row {
			localXMin := true
			switch x {
			case 0:
				localXMin = height < row[x+1]
			case xSize - 1:
				localXMin = prevHeight > height
			default:
				localXMin = prevHeight > height && height < row[x+1]
			}
			prevHeight = height
			if localXMin && localYMin(x, height) {
				localMinima = append(localMinima, point{x, y})
			}
		}
	}
	return localMinima
}
