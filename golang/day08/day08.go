package day08

import (
	"strings"
)

type signalMapping struct {
	inputs  []byte
	outputs []byte
}

func part1(input []string) int {
	mappings := readInput(input)
	count := 0
	for _, m := range mappings {
		for _, o := range m.outputs {
			switch countSetBits(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func part2(input []string) int {
	sumDisplays := 0
	mappings := readInput(input)
	for _, m := range mappings {
		displayedNum := decode(m)
		sumDisplays += displayedNum
	}
	return sumDisplays
}

func decode(m signalMapping) int {
	known := make([]byte, 10)
	others := make([]byte, 0, 6)
	for _, b := range m.inputs {
		switch countSetBits(b) {
		case 2:
			known[1] = b
		case 3:
			known[7] = b
		case 4:
			known[4] = b
		case 7:
			known[8] = b
		default:
			others = append(others, b)
		}
	}
	one := known[1]
	four := known[4]
	for _, b := range others {
		numSet := countSetBits(b)
		common := countSetBits(b&one) + countSetBits(b&four)
		digit := getDigit(numSet, common)
		known[digit] = b
	}
	pow := 1000
	displayedNum := 0
	for _, o := range m.outputs {
		digit := 0
		for ; digit < 10 && o^known[digit] != 0; digit++ {
		}
		displayedNum += digit * pow
		pow /= 10
	}
	return displayedNum
}

func getDigit(numBitsSet, commonWithOneAndFour int) int {
	switch commonWithOneAndFour {
	case 3:
		return 2
	case 6:
		return 9
	case 4:
		if numBitsSet == 5 {
			return 5
		} else {
			return 6
		}
	case 5:
		if numBitsSet == 5 {
			return 3
		} else {
			return 0
		}
	default: // never happens
		return -1
	}
}

func readInput(input []string) []signalMapping {
	mappings := make([]signalMapping, len(input))
	for i, s := range input {
		mappings[i] = makeSignalMapping(s)
	}
	return mappings
}

func makeSignalMapping(s string) signalMapping {
	parts := strings.Split(s, " | ")
	rawInputs := strings.Split(parts[0], " ")
	rawOutputs := strings.Split(parts[1], " ")
	return signalMapping{
		stringsToBytes(rawInputs),
		stringsToBytes(rawOutputs),
	}
}

func stringsToBytes(s []string) []byte {
	result := make([]byte, len(s))
	for i, r := range s {
		result[i] = convertToByte(r)
	}
	return result
}

func convertToByte(s string) byte {
	result := byte(0)
	for _, r := range s {
		result |= 1 << (byte(r) - 0b01100001)
	}
	return result
}

func countSetBits(n byte) (count int) {
	for ; n > 0; n = n & (n - 1) {
		count++
	}
	return
}
