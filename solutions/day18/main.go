package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(18, part1, part2)
}

func part1(
	input string,
) string {
	m := buildMatrix()
	cs, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	for _, c := range cs[:1024] {
		m[c.Y][c.X] = true
	}

	start := util.Coordinate{X: 0, Y: 0}
	end := util.Coordinate{X: 70, Y: 70}
	visitedSet := buildVisitedSet[int]()
	shortestPath(m, start, end, 0, visitedSet)

	shortest := visitedSet[70][70]
	return fmt.Sprintf("Shortest path is %d steps", shortest)
}

func part2(
	input string,
) string {
	m := buildMatrix()
	cs, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	// Initialize everything
	latest := cs[0]
	cs = cs[1:]
	winningSet := buildVisitedSet[bool]()
	m[latest.Y][latest.X] = true
	count := 1
	start := util.Coordinate{X: 0, Y: 0}
	end := util.Coordinate{X: 70, Y: 70}
	success := anyPath(m, start, end, buildVisitedSet[bool](), winningSet)

	for success && len(cs) > 0 {
		for !winningSet[latest.Y][latest.X] && len(cs) > 0 {
			count++
			latest = cs[0]
			cs = cs[1:]
			m[latest.Y][latest.X] = true
		}

		winningSet = buildVisitedSet[bool]()
		success = anyPath(m, start, end, buildVisitedSet[bool](), winningSet)
	}

	return fmt.Sprintf("The byte that broke the camel's back was %s", latest)
}

func buildMatrix() [][]bool {
	out := make([][]bool, 71)
	for y := 0; y <= 70; y++ {
		out[y] = make([]bool, 71)
	}
	return out
}

func buildVisitedSet[T any]() map[int]map[int]T {
	visitedSet := make(map[int]map[int]T)
	for y := 0; y <= 70; y++ {
		visitedSet[y] = make(map[int]T)
	}
	return visitedSet
}

func shortestPath(
	m [][]bool,
	c util.Coordinate,
	goal util.Coordinate,
	steps int,
	visitedSet map[int]map[int]int,
) {
	visit := visitedSet[c.Y][c.X]
	if visit > 0 && visit <= steps {
		// Someone beat us to this position
		return
	}

	endVisit := visitedSet[goal.Y][goal.X]
	if endVisit > 0 && endVisit <= steps {
		// Someone beat us to the end goal
		return
	}

	if !util.In2DArray(c, m) || m[c.Y][c.X] {
		// The current position is impassable or outside the matrix
		return
	}

	// Register our place on the map
	visitedSet[c.Y][c.X] = steps

	// Travel new places
	shortestPath(m, c.North(), goal, steps+1, visitedSet)
	shortestPath(m, c.East(), goal, steps+1, visitedSet)
	shortestPath(m, c.South(), goal, steps+1, visitedSet)
	shortestPath(m, c.West(), goal, steps+1, visitedSet)
}

func anyPath(
	m [][]bool,
	c util.Coordinate,
	goal util.Coordinate,
	visitedSet map[int]map[int]bool,
	winningSet map[int]map[int]bool,
) bool {
	if visitedSet[c.Y][c.X] || !util.In2DArray(c, m) || m[c.Y][c.X] {
		// Been here, outside matrix or position is impassable
		return false
	} else if c.Equals(goal) {
		// We did it!
		winningSet[c.Y][c.X] = true
		return true
	}

	visitedSet[c.Y][c.X] = true
	winner := anyPath(m, c.South(), goal, visitedSet, winningSet) ||
		anyPath(m, c.East(), goal, visitedSet, winningSet) ||
		anyPath(m, c.West(), goal, visitedSet, winningSet) ||
		anyPath(m, c.North(), goal, visitedSet, winningSet)
	if winner {
		winningSet[c.Y][c.X] = true
	}

	return winner
}
