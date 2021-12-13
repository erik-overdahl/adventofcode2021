package day12

import (
	"strings"
)

func part1(input []string) int {
	graph := makeGraph(input)
	memoizer := make(map[int]int)
	result := countPaths(graph, 0, 1<<len(graph.positions), false, memoizer)
	return result
}

func part2(input []string) int {
	graph := makeGraph(input)
	memoizer := make(map[int]int)
	result := countPaths(graph, 0, 1<<len(graph.positions), true, memoizer)
	return result
}

type graph struct {
	positions []string
	edges     [][]int
}

func countPaths(g *graph, v, seen int, canReturn bool, memoizer map[int]int) int {
	if v == 1 { // end
		return 1
	}
	memKey := (seen << 9) + (v << 1)
	if canReturn {
		memKey |= 1
	}
	if result, exists := memoizer[memKey]; exists {
		return result
	}
	visited := bitAt(seen, v) == 1
	if visited {
		if v == 0 { // start
			return 0
		}
		smallCave := byte(g.positions[v][0]) > 90
		if smallCave && !canReturn {
			return 0
		} else if smallCave && canReturn {
			canReturn = false
		}
	}
	seen = setBit(seen, v)
	sum := 0
	for _, n := range g.edges[v] {
		sum += countPaths(g, n, seen, canReturn, memoizer)
	}
	memoizer[memKey] = sum
	return sum
}

func makeGraph(input []string) *graph {
	positions := map[string]int{"start": 0, "end": 1}
	edges := [][]int{{}, {}}
	for _, line := range input {
		nodes := strings.Split(line, "-")
		pos1, exists := positions[nodes[0]]
		if !exists {
			pos1 = len(edges)
			positions[nodes[0]] = pos1
			edges = append(edges, []int{})
		}
		pos2, exists := positions[nodes[1]]
		if !exists {
			pos2 = len(edges)
			positions[nodes[1]] = pos2
			edges = append(edges, []int{})
		}
		edges[pos1] = append(edges[pos1], pos2)
		edges[pos2] = append(edges[pos2], pos1)
	}
	gPositions := make([]string, len(positions))
	for name, pos := range positions {
		gPositions[pos] = name
	}
	return &graph{positions: gPositions, edges: edges}
}

func bitAt(n int, pos int) int {
	return (n >> pos) & 1
}

func setBit(n int, pos int) int {
	return n | 1<<pos
}
