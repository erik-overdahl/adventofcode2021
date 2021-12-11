package day10

import (
	"sort"
)

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
				expected := pair(b)
				if open != expected {
					corrupted = true
					score += getScore(b)
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
				expected := pair(b)
				if open != expected {
					corrupted = true
				}
			}
		}
		if !corrupted {
			score := 0
			for stack.Len() > 0 {
				b := stack.Pop()
				val := getScore(b)
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

func getScore(b byte) int {
	switch b {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		return 0
	}
}

func pair(b byte) byte {
	switch b {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	case '>':
		return '<'
	default:
		return 0
	}
}
