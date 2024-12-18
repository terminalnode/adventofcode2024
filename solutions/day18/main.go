package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(18, part1, nil)
}

func part1(
	input string,
) string {
	out, err := run(input, 70, 70, 1024)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Shortest path is %d steps", out)
}

func run(
	input string,
	maxX int,
	maxY int,
	fallenBytes int,
) (int, error) {
	m := buildMatrix(maxX, maxY)
	cs, err := parse(input)
	if err != nil {
		return -1, err
	}

	if len(cs) > fallenBytes {
		cs = cs[:fallenBytes]
	}

	for _, c := range cs {
		m[c.Y][c.X] = true
	}

	visitedSet := buildVisitedSet(maxY)
	shortestPath(
		m,
		util.Coordinate{X: 0, Y: 0},
		util.Coordinate{X: maxX, Y: maxY},
		0,
		visitedSet,
	)

	return visitedSet[maxY][maxX], nil
}

func buildMatrix(
	maxX int,
	maxY int,
) [][]bool {
	out := make([][]bool, maxY+1)
	for y := 0; y <= maxY; y++ {
		out[y] = make([]bool, maxX+1)
	}
	return out
}

func buildVisitedSet(
	maxY int,
) map[int]map[int]int {
	visitedSet := make(map[int]map[int]int)
	for y := 0; y <= maxY; y++ {
		visitedSet[y] = make(map[int]int)
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
