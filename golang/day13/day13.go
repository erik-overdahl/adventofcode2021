package day13

import (
	// "fmt"
	"fmt"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func part1(input []string) int {
	points, reflections := readInput(input)
	points = applyReflections(points, reflections[:1])
	return len(points)
}

func part2(input []string) int {
	points, reflections := readInput(input)
	points = applyReflections(points, reflections)
	xMax, yMax := 0, 0
	for _, p := range points {
		if p.x > xMax {
			xMax = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
	}
	output := make([][]byte, yMax+1)
	for i := 0; i <= yMax; i++ {
		row := make([]byte, xMax+1)
		for j := 0; j <= xMax; j++ {
			row[j] = ' '
		}
		output[i] = row
	}
	for _, p := range points {
		output[p.y][p.x] = '#'
	}
	s := ""
	for _, row := range output {
		s += "\n" + string(row)
	}
	fmt.Println(s)
	return 0
}

func applyReflections(points []point, reflections [](func(point) point)) []point {
	seen := make([]point, 0, len(points))
	for _, f := range reflections {
		for _, p := range points {
			p = f(p)
			exists := false
			for _, pt := range seen {
				if (p.x == pt.x) && (p.y == pt.y) {
					exists = true
				}
			}
			if !exists {
				seen = append(seen, p)
			}
		}
		points = seen
		seen = []point{}
	}
	return points
}

func readInput(input []string) ([]point, [](func(point) point)) {
	size := len(input)
	points := []point{}
	reflections := [](func(point) point){}
	nums := make([]int, 2)
	i := 0
	for ; input[i] != ""; i++ {
		vals := strings.Split(input[i], ",")
		for j, s := range vals {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums[j] = n
		}
		points = append(points, point{nums[0], nums[1]})
	}
	i++
	for ; i < size; i++ {
		parts := strings.Split(input[i], " ")
		r := []byte(parts[2])
		n, err := strconv.Atoi(string(r[2:]))
		if err != nil {
			panic(err)
		}
		var f func(point) point
		if r[0] == 'x' {
			f = func(p point) point {
				return reflectX(p, n)
			}
		} else if r[0] == 'y' {
			f = func(p point) point {
				return reflectY(p, n)
			}
		}
		reflections = append(reflections, f)
	}
	return points, reflections
}

func reflectX(p point, axis int) point {
	if p.x > axis {
		p.x -= 2 * (p.x - axis)
	}
	return p
}

func reflectY(p point, axis int) point {
	if p.y > axis {
		p.y -= 2 * (p.y - axis)
	}
	return p
}
