package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Graph struct {
	nodes         map[string]int
	names         map[int]string
	adjacencyList [][]int
	revisitable   map[int]bool
}

func (g *Graph) print() {
	for n, edges := range g.adjacencyList {
		fmt.Printf("node: %v, edges: ", g.names[n])
		for _, to := range edges {
			fmt.Printf("%v, ", g.names[to])
		}
		fmt.Println()
	}
}

func IsUpperCase(s string) bool {
	return s == strings.ToUpper(s)
}

func IsAlreadyOnPath(n int, path []int) bool {
	for _, p := range path {
		if n == p {
			return true
		}
	}
	return false
}

func (g *Graph) findAllPaths(from, to int, withOneRevisit bool) [][]int {
	allPaths := [][]int{}
	beingVisited := make([]int, len(g.adjacencyList))

	var dfs func(node int, path []int, revisitAllowed bool)
	dfs = func(node int, path []int, revisitAllowed bool) {
		beingVisited[node]++

		if node == to {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
		}

		for _, nextNode := range g.adjacencyList[node] {
			if (nextNode != from && nextNode != to) && revisitAllowed && !g.revisitable[nextNode] && IsAlreadyOnPath(nextNode, path) {
				path = append(path, nextNode)
				dfs(nextNode, path, false)
				path = path[:len(path)-1]
			} else if g.revisitable[nextNode] || beingVisited[nextNode] == 0 {
				path = append(path, nextNode)
				dfs(nextNode, path, revisitAllowed)
				path = path[:len(path)-1]
			}
		}

		beingVisited[node]--
	}

	path := []int{g.nodes["start"]}
	dfs(from, path, withOneRevisit)

	return allPaths
}

func parseIngput(lines []string) *Graph {
	g := &Graph{
		nodes:         map[string]int{},
		names:         map[int]string{},
		adjacencyList: [][]int{},
		revisitable:   map[int]bool{},
	}

	node := 0
	r := regexp.MustCompile(`^(\w+)?-(\w+)$`)

	// parse input
	for _, line := range lines {
		groups := r.FindStringSubmatch(line)
		from, to := groups[1], groups[2]

		// check if 'from'-node already exists
		fromId, ok := g.nodes[from]
		if !ok {
			g.nodes[from] = node
			g.names[node] = from
			g.adjacencyList = append(g.adjacencyList, make([]int, 0))
			g.revisitable[node] = IsUpperCase(from)
			fromId = node
			node++
		}

		// check if 'to'-node already exists
		toId, ok := g.nodes[to]
		if !ok {
			g.nodes[to] = node
			g.names[node] = to
			g.adjacencyList = append(g.adjacencyList, make([]int, 0))
			g.revisitable[node] = IsUpperCase(to)
			toId = node
			node++
		}

		// add two edges to the adjacency list
		g.adjacencyList[fromId] = append(g.adjacencyList[fromId], toId)
		g.adjacencyList[toId] = append(g.adjacencyList[toId], fromId)
	}

	return g
}

func partOne(lines []string) int {
	g := parseIngput(lines)
	allPaths := g.findAllPaths(g.nodes["start"], g.nodes["end"], false)
	return len(allPaths)
}

func partTwo(lines []string) int {
	g := parseIngput(lines)
	allPaths := g.findAllPaths(g.nodes["start"], g.nodes["end"], true)
	return len(allPaths)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
