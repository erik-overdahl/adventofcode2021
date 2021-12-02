package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarinePosition struct {
	horiz int
	depth int
}

func Part1(input []string) int {
	pos := SubmarinePosition{0, 0}
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
		case "down":
			pos.depth += magnitude
		case "up":
			pos.depth -= magnitude
		}
	}
	return pos.horiz * pos.depth
}
