package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/juju/collections/deque"
	"github.com/nobbs/advent-of-code-in-go/util"
)

const startBag string = "shiny gold bags"

func bfs(adjacency [][]int, start int) int {
	d := deque.New()
	visited := make([]bool, len(adjacency))
	result := 0

	d.PushBack(start)
	visited[start] = true

	for d.Len() > 0 {
		next, _ := d.PopFront()

		for i, v := range adjacency[next.(int)] {
			if v > 0 && !visited[i] {
				result++
				d.PushBack(i)
				visited[i] = true
			}
		}
	}

	return result
}

func dfs(adjacency [][]int, start int, visited []bool) int {
	visited[start] = true

	sum := 1

	for i, v := range adjacency[start] {
		if v > 0 {
			sum += v * dfs(adjacency, i, visited)
		}
	}

	return sum
}

func partOne(lines []string) int {
	numBags := len(lines)
	bagNames := make([]string, numBags)

	adjacency := make([][]int, numBags)

	findBagIndex := func(bag string) int {
		for i, v := range bagNames {
			if v == bag {
				return i
			}
		}
		return -1
	}

	for i := range adjacency {
		adjacency[i] = make([]int, numBags)
	}

	for i, line := range lines {
		bagNames[i] = regexp.MustCompile(`^([\w\s]+) contain`).FindStringSubmatch(line)[1]
	}

	for i, line := range lines {
		children := regexp.MustCompile(`(\d+) ([\w\s]*)[,\.]`).FindAllStringSubmatch(line, -1)
		for _, c := range children {
			child := c[2]
			n := util.ParseInt(c[1])

			if !strings.HasSuffix(child, "s") {
				child = child + string("s")
			}

			childIndex := findBagIndex(child)
			adjacency[childIndex][i] = n
		}
	}

	startIndex := findBagIndex(startBag)
	return bfs(adjacency, startIndex)
}

func partTwo(lines []string) int {
	numBags := len(lines)
	bagNames := make([]string, numBags)

	adjacency := make([][]int, numBags)

	findBagIndex := func(bag string) int {
		for i, v := range bagNames {
			if v == bag {
				return i
			}
		}
		return -1
	}

	for i := range adjacency {
		adjacency[i] = make([]int, numBags)
	}

	for i, line := range lines {
		bagNames[i] = regexp.MustCompile(`^([\w\s]+) contain`).FindStringSubmatch(line)[1]
	}

	for i, line := range lines {
		children := regexp.MustCompile(`(\d+) ([\w\s]*)[,\.]`).FindAllStringSubmatch(line, -1)
		for _, c := range children {
			child := c[2]
			n := util.ParseInt(c[1])

			if !strings.HasSuffix(child, "s") {
				child = child + string("s")
			}

			childIndex := findBagIndex(child)
			adjacency[i][childIndex] = n
		}
	}

	startIndex := findBagIndex(startBag)
	visited := make([]bool, len(adjacency))
	visited[startIndex] = true

	return dfs(adjacency, startIndex, visited) - 1
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
