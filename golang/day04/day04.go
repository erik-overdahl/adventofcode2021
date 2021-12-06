package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type BingoCard struct {
	nums   [25]int
	marked [25]bool
}

func (b *BingoCard) ToString() string {
	red := "\033[31m"
	colorOff := "\033[0m"
	s := ""
	for i := 0; i < 25; i += 5 {
		row := ""
		for j := 0; j < 5; j++ {
			cell := b.nums[i+j]
			sCell := fmt.Sprintf("%2d", cell)
			if b.marked[i+j] {
				sCell = fmt.Sprintf("%s%s%s", red, sCell, colorOff)
			}
			row = fmt.Sprintf("%s %s", row, sCell)
		}
		s = fmt.Sprintf("%s%s\n", s, row)
	}
	return s
}

func Part1(input []string) int {
	draws, cards := readInput(input)
	for _, draw := range draws {
		for _, card := range cards {
			marked := MarkIfPresent(card, draw)
			if marked && (CheckCols(card) || CheckRows(card)) {
				return SumUnmarked(card) * draw
			}
		}
	}
	return 0
}

func Part2(input []string) int {
	draws, cards := readInput(input)
	size := len(draws)
	lastWinningDraw := 0
	var lastWinningCard *BingoCard
	for _, card := range cards {
		for i := 0; i < size; i++ {
			marked := MarkIfPresent(card, draws[i])
			if marked && (CheckCols(card) || CheckRows(card)) {
				if i > lastWinningDraw {
					lastWinningDraw = i
					lastWinningCard = card
				}
				break
			}
		}
	}
	return SumUnmarked(lastWinningCard) * draws[lastWinningDraw]
}

func NewBingoCard(lines []string) *BingoCard {
	nums := [25]int{}
	for i, line := range lines {
		row := i * 5
		c := 0
		for _, s := range strings.Split(line, " ") {
			if len(s) > 0 {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums[row+c] = n
				c++
			}
		}
	}
	return &BingoCard{nums: nums, marked: [25]bool{}}
}

func MarkIfPresent(card *BingoCard, draw int) bool {
	for i, n := range card.nums {
		if n == draw {
			card.marked[i] = true
			return true
		}
	}
	return false
}

func CheckRows(b *BingoCard) bool {
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

func CheckCols(b *BingoCard) bool {
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

func SumUnmarked(b *BingoCard) int {
	sum := 0
	for i, marked := range b.marked {
		if !marked {
			sum += int(b.nums[i])
		}
	}
	return sum
}

func readDraws(rawline string) []int {
	var draws []int
	for _, s := range strings.Split(rawline, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		draws = append(draws, n)
	}
	return draws
}

func readInput(input []string) (draws []int, cards []*BingoCard) {
	inputSize := len(input)
	draws = readDraws(input[0])
	for i := 2; i < inputSize; i += 6 {
		card := NewBingoCard(input[i : i+6])
		cards = append(cards, card)
	}
	return
}
