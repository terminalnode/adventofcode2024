package main

import (
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strings"
)

// Due to function types not implementing == or !=, I need to store the direction as an int.

type Coordinate = util.Coordinate
type BoolMatrix = util.BoolMatrix
type direction = int

var directions = []util.Direction{
	Coordinate.North,
	Coordinate.East,
	Coordinate.South,
	Coordinate.West,
}

type parsedInput struct {
	Guard       Coordinate
	Direction   direction
	ObstacleMap BoolMatrix
	VisitedMap  BoolMatrix
}

func main() {
	common.Setup(6, part1, nil)
}

func part1(
	input string,
) string {
	parsed, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	for parsed.move() {
		// Do nothing, move() does all the work
	}
	count := util.CountInMatrix(parsed.VisitedMap, true)
	return fmt.Sprintf("Visited spots in the Matrix: %d", count)
}

func parseInput(
	input string,
) (parsedInput, error) {
	var guard Coordinate
	var obstacles BoolMatrix
	var visited BoolMatrix
	var err error

	lines := strings.Split(input, "\n")
	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	rawObstacles := make([][]bool, maxY+1)
	rawVisited := make([][]bool, maxY+1)

	foundGuard := false
	for y := 0; y <= maxY; y++ {
		rawObstacles[y] = make([]bool, maxX+1)
		rawVisited[y] = make([]bool, maxX+1)

		for x := 0; x <= maxX; x++ {
			c := lines[y][x]
			rawObstacles[y][x] = c == '#'
			if c == '^' {
				foundGuard = true
				guard = Coordinate{X: x, Y: y}
				rawVisited[y][x] = true
			}
		}
	}
	if foundGuard == false {
		return parsedInput{}, errors.New("could not find guard in matrix")
	}

	obstacles, err = util.NewMatrixFromRows(rawObstacles)
	if err != nil {
		return parsedInput{}, fmt.Errorf("failed to create obstacle matrix: %v", err)
	}

	visited, err = util.NewMatrixFromRows(rawVisited)
	if err != nil {
		return parsedInput{}, fmt.Errorf("failed to create visited matrix: %v", err)
	}

	return parsedInput{
		Guard:       guard,
		Direction:   0, // This means north
		ObstacleMap: obstacles,
		VisitedMap:  visited,
	}, nil
}

// Move and return a boolean indicating whether the guard is inside the matrix or not
func (
	p *parsedInput,
) move() bool {
	newPos := getNewPosition(p.Guard, p.Direction)
	isBlocked, err := p.ObstacleMap.Get(newPos.X, newPos.Y)
	if err != nil {
		return false
	} else if isBlocked {
		p.rotate()
		newPos = getNewPosition(p.Guard, p.Direction)
	}
	p.Guard = newPos

	err = p.VisitedMap.Set(p.Guard.X, p.Guard.Y, true)
	return err == nil
}

func (p *parsedInput) rotate() {
	p.Direction = (p.Direction + 1) % 4
}

func getNewPosition(
	guard Coordinate,
	direction int,
) Coordinate {
	dirF := directions[direction]
	return dirF(guard)
}
