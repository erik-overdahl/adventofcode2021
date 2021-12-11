package day05

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func makePoint(x, y int) Point {
	return Point{x, y}
}

func makeLine(p1, p2 Point) Line {
	// always put leftmost/uppermost point first
	if (p1.x < p2.x) || ((p1.x == p2.x) && (p1.y > p2.y)) {
		return Line{p1, p2}
	} else {
		return Line{p2, p1}
	}
}

func stringToPoint(s string) Point {
	parts := strings.Split(s, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return makePoint(x, y)
}

func readInput(input []string) []Line {
	lines := make([]Line, len(input))
	for i, line := range input {
		pts := strings.Split(line, " -> ")
		p1 := stringToPoint(pts[0])
		p2 := stringToPoint(pts[1])
		lines[i] = makeLine(p1, p2)
	}
	return lines
}

func boundaries(lines []Line) (xMax, yMax int) {
	for _, l := range lines {
		if l.end.x > xMax {
			xMax = l.end.x
		}
		if l.start.y > yMax {
			yMax = l.start.y
		} else if l.end.y > yMax {
			yMax = l.end.y
		}
	}
	xMax++
	yMax++
	return
}

func pointsCoveredVerticalHorizontal(lines []Line) []byte {
	xMax, yMax := boundaries(lines)
	covered := make([]byte, (xMax*yMax)+1)
	for _, line := range lines {
		p1 := line.start
		p2 := line.end
		if p1.x == p2.x { // vertical, p1 on top
			x := p1.x
			for y := p1.y; y >= p2.y; y-- {
				covered[(y*xMax)+x]++
			}
		} else if p1.y == p2.y { // horizontal, p1 on left
			y := p1.y
			for x := p1.x; x <= p2.x; x++ {
				covered[(y*xMax)+x]++
			}
		}
	}
	return covered
}

func pointsCovered(lines []Line) []byte {
	xMax, _ := boundaries(lines)
	covered := pointsCoveredVerticalHorizontal(lines)
	for _, line := range lines {
		p1 := line.start
		p2 := line.end
		// only consider diagonal lines; p1 always leftmost
		if (p1.x != p2.x) && (p1.y < p2.y) {
			for x, y := p1.x, p1.y; x <= p2.x && y <= p2.y; x, y = x+1, y+1 {
				covered[(y*xMax)+x]++
			}
		} else if (p1.x != p2.x) && (p1.y > p2.y) {
			for x, y := p1.x, p1.y; x <= p2.x && y >= p2.y; x, y = x+1, y-1 {
				covered[(y*xMax)+x]++
			}

		}
	}
	return covered
}

func numPointsOverlapped(covered []byte) int {
	total := 0
	for _, n := range covered {
		if n > 1 {
			total++
		}
	}
	return total
}

func part1(lines []Line) int {
	visited := pointsCoveredVerticalHorizontal(lines)
	return numPointsOverlapped(visited)
}

func part2(lines []Line) int {
	visited := pointsCovered(lines)
	return numPointsOverlapped(visited)
}
