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

// VisitedDirectionMap is meant to go from x => y => direction
type tVisitedMap map[int]map[int]map[int]bool

var directions = []util.Direction{
	Coordinate.North,
	Coordinate.East,
	Coordinate.South,
	Coordinate.West,
}

type parsedInput struct {
	Guard          Coordinate
	Direction      direction
	ObstacleMatrix BoolMatrix
	VisitedMap     tVisitedMap
}

func main() {
	common.Setup(6, part1, part2)
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

	var count int
	for _, xMap := range parsed.VisitedMap {
		count += len(xMap)
	}

	return fmt.Sprintf("Visited spots in the Matrix: %d", count)
}

func part2(
	input string,
) string {
	// Solve part 1 first to know where to place obstacles
	original, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	originalX := original.Guard.X
	originalY := original.Guard.Y
	for original.move() {
		// Do nothing, move() does all the work
	}

	// Then loop through all positions the guard will visit and place an obstacle
	count := 0
	attempts := 0
	//for x := 0; x <= original.ObstacleMatrix.MaxX; x++ {
	for x, yMap := range original.VisitedMap {
		//for y := 0; y <= original.ObstacleMatrix.MaxY; y++ {
		for y := range yMap {
			if y == originalY && x == originalX {
				continue
			}
			attempts++

			// Reset the map with the new obstacle
			original.Guard = Coordinate{X: originalX, Y: originalY}
			original.Direction = 0
			original.VisitedMap = make(tVisitedMap)
			_ = original.ObstacleMatrix.Set(x, y, true)

			// Run this variant
			for original.move() {
				// Do nothing, move() does all the work
			}
			_ = original.ObstacleMatrix.Set(x, y, false)

			// If the guard is still in the matrix, he's in a loop
			if original.guardIsInMatrix() {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d of the available %d positions make him loopy!", count, attempts)
}

func parseInput(
	input string,
) (parsedInput, error) {
	var guard Coordinate
	var obstacles BoolMatrix
	var err error

	lines := strings.Split(input, "\n")
	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	rawObstacles := make([][]bool, maxY+1)
	visitedMap := make(tVisitedMap)

	foundGuard := false
	for y := 0; y <= maxY; y++ {
		rawObstacles[y] = make([]bool, maxX+1)

		for x := 0; x <= maxX; x++ {
			c := lines[y][x]
			rawObstacles[y][x] = c == '#'
			if c == '^' {
				foundGuard = true
				guard = Coordinate{X: x, Y: y}
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

	out := parsedInput{
		Guard:          guard,
		Direction:      0, // This means north
		ObstacleMatrix: obstacles,
		VisitedMap:     visitedMap,
	}
	out.checkIfVisitedAndMark(guard)

	return out, nil
}

// Move and return a boolean indicating whether the guard is stuck in a loop or inside the matrix
// True means he's stuck or out, false means he's inside on a non-visited spot.
func (p *parsedInput) move() bool {
	newPos := getNewPosition(p.Guard, p.Direction)
	isBlocked, err := p.ObstacleMatrix.Get(newPos.X, newPos.Y)
	if err != nil {
		p.Guard = newPos
		return false
	} else if isBlocked {
		p.rotate()
		return !p.checkIfVisitedAndMark(p.Guard)
	}
	p.Guard = newPos

	if p.checkIfVisitedAndMark(newPos) {
		return false
	}

	return true
}

func (p *parsedInput) guardIsInMatrix() bool {
	return p.ObstacleMatrix.IsInMatrix(p.Guard.X, p.Guard.Y)
}

func (p *parsedInput) checkIfVisitedAndMark(c Coordinate) bool {
	if p.VisitedMap[c.X] == nil {
		p.VisitedMap[c.X] = make(map[int]map[int]bool)
	}
	if p.VisitedMap[c.X][c.Y] == nil {
		p.VisitedMap[c.X][c.Y] = make(map[int]bool)
	}
	old := p.VisitedMap[c.X][c.Y][p.Direction]
	p.VisitedMap[c.X][c.Y][p.Direction] = true
	return old
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
