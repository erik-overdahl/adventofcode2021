package day12

import (
	"strings"
)

type graph struct {
	adjList  map[string][]string
	visited  map[string]bool
	numPaths int
}

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
	visited := make(map[string]bool)
	for k := range adjacencyList {
		visited[k] = false
	}
	return &graph{adjList: adjacencyList, visited: visited, numPaths: 0}
}

func (g *graph) countPaths(v string) {
	if v == "end" {
		g.numPaths++
	} else {
		if byte(v[0]) > 90 {
			g.visited[v] = true
		}
		for _, n := range g.adjList[v] {
			if !g.visited[n] {
				g.countPaths(n)
			}
		}
		g.visited[v] = false
	}
}

func part1(input []string) int {
	graph := MakeGraph(input)
	graph.countPaths("start")
	return graph.numPaths
}
