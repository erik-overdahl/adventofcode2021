package utils

type AOCDay interface {
	Day() int
	Init(input string)
	Part1() int
	Part2() int
}
