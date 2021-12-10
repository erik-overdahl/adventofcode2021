package day10

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
		// if the line is incomplete, skip for now
		// if stack.Len() > 0 {
		// 	fmt.Printf("Line incomplete!\n")
		// }
	}
	return score
}
