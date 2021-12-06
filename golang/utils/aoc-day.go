package utils

type AOCDay interface {
	Day() int
	Init(input string)
	Part1() string
	Part2() string
}
