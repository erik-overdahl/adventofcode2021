package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarinePosition struct {
	aim   int
	horiz int
	depth int
}

func Part1(input []string) int {
	pos := calculatePosition(input)
	return pos.horiz * pos.aim
}

func Part2(input []string) int {
	pos := calculatePosition(input)
	return pos.horiz * pos.depth
}

func calculatePosition(input []string) SubmarinePosition {
	pos := SubmarinePosition{0, 0, 0}
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Errorf("Unable to convert '%s' to int: %v", parts[1], err))
		}
		switch direction {
		case "forward":
			pos.horiz += magnitude
			pos.depth += magnitude * pos.aim
		case "down":
			pos.aim += magnitude
		case "up":
			pos.aim -= magnitude
		}
	}
	return pos
}
