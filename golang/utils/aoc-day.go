package utils

type AOCDay interface {
	Day() int
	Init(filename string)
	Part1() string
	Part2() string
}
