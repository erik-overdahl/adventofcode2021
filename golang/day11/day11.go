package day11

type point struct{ x, y int }

type stack struct {
	stack []point
	ptr   int
}

func MakeStack() *stack {
	s := make([]point, 0)
	return &stack{s, -1}
}

func (s *stack) Len() int {
	return len(s.stack)
}

func (s *stack) Push(p point) {
	s.stack = append(s.stack, p)
	s.ptr++
}

func (s *stack) Pop() point {
	p := s.stack[s.ptr]
	s.stack = s.stack[:s.ptr]
	s.ptr--
	return p
}

type board struct {
	Current [][]byte
	Next    [][]byte
	stack   *stack
}

func makeBoard(start [][]byte) *board {
	internal := make([][]byte, len(start))
	Next := make([][]byte, len(start))
	for i, row := range start {
		internal[i] = make([]byte, len(row))
		for j, e := range row {
			internal[i][j] = e
		}
		Next[i] = make([]byte, len(row))
	}
	return &board{Current: internal, Next: Next, stack: MakeStack()}
}

func (b *board) NextGen() {
	for y, row := range b.Current {
		for x, energy := range row {
			if energy == 9 {
				b.stack.Push(point{x, y})
				b.Next[y][x] = 0
			} else {
				b.Next[y][x] = energy + 1
			}
		}
	}
	for b.stack.Len() > 0 {
		p := b.stack.Pop()
		neighbors := b.NeighborsOf(p)
		for _, n := range neighbors {
			energy := b.Next[n.y][n.x]
			if energy != 0 {
				energy++
				if energy > 9 {
					b.stack.Push(n)
					energy = 0
				}
				b.Next[n.y][n.x] = energy
			}
		}
	}
	b.Next, b.Current = b.Current, b.Next
}

func (b *board) NeighborsOf(p point) []point {
	neighbors := make([]point, 0, 9)
	yMax, xMax := len(b.Next)-1, len(b.Next[0])-1
	nXMin, nXMax := max(p.x-1, 0), min(p.x+1, xMax)
	nYMin, nYMax := max(p.y-1, 0), min(p.y+1, yMax)
	for y := nYMin; y <= nYMax; y++ {
		for x := nXMin; x <= nXMax; x++ {
			neighbors = append(neighbors, point{x, y})
		}
	}
	return neighbors
}

func part1(octopi [][]byte, numGenerations int) int {
	board := makeBoard(octopi)
	numFlashes := 0
	for i := 0; i < numGenerations; i++ {
		board.NextGen()
		numFlashes += countFlashes(board)
	}
	return numFlashes
}

func part2(octopi [][]byte) int {
	board := makeBoard(octopi)
	target := len(octopi) * len(octopi[0])
	steps := 0
	for countFlashes(board) != target {
		board.NextGen()
		steps++
	}
	return steps
}

func countFlashes(b *board) int {
	numFlashes := 0
	for _, row := range b.Current {
		for _, e := range row {
			if e == 0 {
				numFlashes++
			}
		}
	}
	return numFlashes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
