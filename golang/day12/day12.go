package day12

import (
	"fmt"
	"strings"
)

type graph struct {
	adjList map[string][]string
}

var memoizer = make(map[string]int)

func MakeGraph(input []string) *graph {
	adjacencyList := make(map[string][]string)
	for _, line := range input {
		nodes := strings.Split(line, "-")
		if adjacent, exists := adjacencyList[nodes[0]]; exists {
			adjacencyList[nodes[0]] = append(adjacent, nodes[1])
		} else {
			adjacencyList[nodes[0]] = []string{nodes[1]}
		}
		if adjacent, exists := adjacencyList[nodes[1]]; exists {
			adjacencyList[nodes[1]] = append(adjacent, nodes[0])
		} else {
			adjacencyList[nodes[1]] = []string{nodes[0]}
		}
	}
	adjacencyList["end"] = []string{}
	return &graph{
		adjList: adjacencyList,
	}
}

func (g *graph) countPaths(v string, seen map[string]bool, canReturn bool) int {
	visited := seen[v]
	memKey := fmt.Sprintf("%s_%v_%v", v, seen, canReturn)
	if result, exists := memoizer[memKey]; exists {
		return result
	}
	if v == "end" {
		return 1
	} else if visited {
		if v == "start" {
			return 0
		} else if byte(v[0]) > 90 && !canReturn {
			return 0
		} else if byte(v[0]) > 90 && canReturn {
			canReturn = false
		}
	}
	// fucking fine
	seenCopy := make(map[string]bool)
	for k, v := range seen {
		seenCopy[k] = v
	}
	seenCopy[v] = true
	sum := 0
	for _, n := range g.adjList[v] {
		sum += g.countPaths(n, seenCopy, canReturn)
	}
	memoizer[memKey] = sum
	return sum
}

func part1(input []string) int {
	graph := MakeGraph(input)
	seen := make(map[string]bool)
	for v := range graph.adjList {
		seen[v] = false
	}
	return graph.countPaths("start", seen, false)
}

func part2(input []string) int {
	graph := MakeGraph(input)
	seen := make(map[string]bool)
	for v := range graph.adjList {
		seen[v] = false
	}
	return graph.countPaths("start", seen, true)
}
