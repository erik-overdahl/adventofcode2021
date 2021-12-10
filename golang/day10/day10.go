package day10

import (
	"sort"
)

var pairs = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}

var scores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

type stack struct {
	stack []byte
	ptr   int
}

func MakeStack() *stack {
	s := make([]byte, 0)
	return &stack{s, -1}
}

func (s *stack) Len() int {
	return len(s.stack)
}

func (s *stack) Push(b byte) {
	s.stack = append(s.stack, b)
	s.ptr++
}

func (s *stack) Pop() byte {
	b := s.stack[s.ptr]
	s.stack = s.stack[:s.ptr]
	s.ptr--
	return b
}

func part1(input [][]byte) int {
	score := 0
	for _, line := range input {
		stack := MakeStack()
		lineSize := len(line)
		corrupted := false
		for i := 0; i < lineSize && !corrupted; i++ {
			b := line[i]
			switch b {
			case '(', '{', '[', '<':
				stack.Push(b)
			default:
				open := stack.Pop()
				expected, _ := pairs[b]
				if open != expected {
					corrupted = true
					score += scores[b]
				}
			}
		}
	}
	return score
}

func part2(input [][]byte) int {
	lineScores := []int{}
	for _, line := range input {
		stack := MakeStack()
		lineSize := len(line)
		corrupted := false
		for i := 0; i < lineSize && !corrupted; i++ {
			b := line[i]
			switch b {
			case '(', '{', '[', '<':
				stack.Push(b)
			default:
				open := stack.Pop()
				expected, _ := pairs[b]
				if open != expected {
					corrupted = true
				}
			}
		}
		if !corrupted {
			score := 0
			for stack.Len() > 0 {
				b := stack.Pop()
				val, _ := scores[b]
				score = (score * 5) + val
			}
			lineScores = append(lineScores, score)
		}
	}
	score := getMiddleScore(lineScores)
	return score
}

func getMiddleScore(lineScores []int) int {
	sort.Ints(lineScores)
	pos := len(lineScores) / 2
	return lineScores[pos]
}
