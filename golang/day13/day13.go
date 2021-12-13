package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func part1(input []string) int {
	points, reflections := readInput(input)
	points = distinct(applyReflections(points, reflections[:1]))
	return len(points)
}

func part2(input []string) int {
	points, reflections := readInput(input)
	points = distinct(applyReflections(points, reflections))
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

func applyReflections(points []point, reflections [](func([]point) []point)) []point {
	for _, f := range reflections {
		points = f(points)
	}
	return points
}

func distinct(points []point) []point {
	seen := make([]point, 0, len(points))
	for _, p := range points {
		exists := false
		for i := 0; i < len(seen) && !exists; i++ {
			pt := seen[i]
			if p.x == pt.x && p.y == pt.y {
				exists = true
			}
		}
		if !exists {
			seen = append(seen, p)
		}
	}
	return seen
}

func reflectX(p point, pos int) point {
	if p.x > pos {
		p.x -= 2 * (p.x - pos)
	}
	return p
}

func reflectY(p point, pos int) point {
	if p.y > pos {
		p.y -= 2 * (p.y - pos)
	}
	return p
}

func getReflection(axis byte, pos int) func([]point) []point {
	var f (func(point, int) point)
	switch axis {
	case 'x':
		f = reflectX
	case 'y':
		f = reflectY
	}
	return func(points []point) []point {
		for i, p := range points {
			points[i] = f(p, pos)
		}
		return points
	}
}

func readInput(input []string) ([]point, [](func([]point) []point)) {
	size := len(input)
	points := []point{}
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
	reflections := [](func([]point) []point){}
	for ; i < size; i++ {
		parts := strings.Split(input[i], " ")
		r := []byte(parts[2])
		n, err := strconv.Atoi(string(r[2:]))
		if err != nil {
			panic(err)
		}
		reflections = append(reflections, getReflection(r[0], n))
	}
	return points, reflections
}
