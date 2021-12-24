package day14

// maintain mapping of pair -> (pair, pair)
// e.g. NB -> C would be a mapping NB -> NC, CB
// maintain a count of how many times each has been seen overall
// and a set of pairs currently in the chain
//
// so we want MAPPINGS
// and CURRENT CHAIN
// and TOTALS
// for each round, go over the current chain
//   look up the mappings of each pair
//   add the count of this pair to the counts of the mapped pairs for the next chain
//   and to the totals
// at the end, go over the pairs and split them into their letters

func readInitial(input string) (pairs map[string]int, letterCounts map[byte]int) {
	size := len(input)
	pairs = make(map[string]int)
	letterCounts = make(map[byte]int)
	letters := []byte(input)
	letterCounts[letters[0]] = 1
	for i := 1; i < size; i++ {
		pair := string([]byte{letters[i-1], letters[i]})
		if _, exists := pairs[pair]; exists {
			pairs[pair]++
		} else {
			pairs[pair] = 1
		}
		if _, exists := letterCounts[letters[i]]; exists {
			letterCounts[letters[i]]++
		} else {
			letterCounts[letters[i]] = 1
		}
	}
	return
}

func readMappings(input []string) map[string][]string {
	mappings := make(map[string][]string)
	size := len(input)
	for i := 0; i < size; i++ {
		line := input[i]
		l1, l2, l3 := line[0], line[1], line[6]
		p1 := string([]byte{l1, l2})
		p2 := string([]byte{l1, l3})
		p3 := string([]byte{l3, l2})
		mappings[p1] = []string{string(l3), p2, p3}
	}
	return mappings
}

func length(totals map[byte]int) int {
	l := 0
	for _, n := range totals {
		l += n
	}
	return l
}

func runRound(counts map[string]int, mappings map[string][]string, totals map[byte]int) map[string]int {
	result := make(map[string]int)
	for pair, count := range counts {
		children := mappings[pair]
		letter := byte(children[0][0])
		if _, exists := totals[letter]; !exists {
			totals[letter] = count
		} else {
			totals[letter] += count
		}
		for _, child := range children[1:] {
			if _, exists := result[child]; !exists {
				result[child] = count
			} else {
				result[child] += count
			}
		}
	}
	return result
}

func mostAndLeastCommonDiff(totals map[byte]int) int {
	max, min := 0, 1<<62
	for _, occurences := range totals {
		if occurences > max {
			max = occurences
		} else if occurences < min {
			min = occurences
		}
	}
	return max - min
}

func part1(input []string) int {
	currentPairs, letterCounts := readInitial(input[0])
	mappings := readMappings(input[2:])
	for round := 0; round < 10; round++ {
		currentPairs = runRound(currentPairs, mappings, letterCounts)
	}
	return mostAndLeastCommonDiff(letterCounts)
}

func part2(input []string) int {
	currentPairs, letterCounts := readInitial(input[0])
	mappings := readMappings(input[2:])
	for round := 0; round < 40; round++ {
		currentPairs = runRound(currentPairs, mappings, letterCounts)
	}
	return mostAndLeastCommonDiff(letterCounts)
}
