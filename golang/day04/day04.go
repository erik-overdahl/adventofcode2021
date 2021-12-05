package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type BingoCard struct {
	nums   [25]int8
	marked [25]bool
}

func NewBingoCard(lines []string) *BingoCard {
	nums := [25]int8{}
	for i, line := range lines {
		row := i * 5
		c := 0
		for _, s := range strings.Split(line, " ") {
			if len(s) > 0 {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums[row+c] = int8(n)
				c++
			}
		}
	}
	return &BingoCard{nums: nums, marked: [25]bool{}}
}

func (b *BingoCard) markIfPresent(draw int8) bool {
	for i, n := range b.nums {
		if n == draw {
			b.marked[i] = true
			return true
		}
	}
	return false
}

func (b *BingoCard) CheckRows() bool {
	for i := 0; i < 25; i += 5 {
		rowMarked := true
		for j := 0; j < 5; j++ {
			rowMarked = rowMarked && b.marked[i+j]
		}
		if rowMarked {
			return true
		}
	}
	return false
}

func (b *BingoCard) CheckCols() bool {
	for i := 0; i < 5; i++ {
		colMarked := true
		for j := 0; j < 25; j += 5 {
			colMarked = colMarked && b.marked[i+j]
		}
		if colMarked {
			return true
		}
	}
	return false
}

func (b *BingoCard) SumUnmarked() int {
	sum := 0
	for i, marked := range b.marked {
		if marked {
			sum += int(b.nums[i])
		}
	}
	return sum
}

func (b *BingoCard) ToString() string {
	bold := "\033[1m"
	s := ""
	for i := 0; i < 25; i += 5 {
		row := ""
		for j := 0; j < 5; j++ {
			cell := b.nums[i+j]
			sCell := fmt.Sprintf("%d", cell)
			if b.marked[i+j] {
				sCell = fmt.Sprintf("%s%s", bold, sCell)
			}
			row = fmt.Sprintf("%s %s", row, sCell)
		}
		s = fmt.Sprintf("%s%s\n", s, row)
	}
	return s
}

func Part1(input []string) int {
	inputSize := len(input)
	draws := readDraws(input[0])
	var cards []*BingoCard
	for i := 2; i < inputSize; i += 6 {
		card := NewBingoCard(input[i : i+6])
		cards = append(cards, card)
	}
	for _, draw := range draws {
		for _, card := range cards {
			marked := card.markIfPresent(draw)
			if marked {
				fmt.Printf("%s\n", card.ToString())
				if card.CheckCols() || card.CheckRows() {
					return card.SumUnmarked() * int(draw)
				}
			}
		}
	}
	return 0
}

func readDraws(rawline string) []int8 {
	var draws []int8
	for _, s := range strings.Split(rawline, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		draws = append(draws, int8(n))
	}
	return draws
}
