package day08

import "strings"

func part1(input []string) int {
	count := 0
	for _, s := range input {
		parts := strings.Split(s, " | ")
		outputs := strings.Split(parts[1], " ")
		for _, o := range outputs {
			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}
