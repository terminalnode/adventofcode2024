package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"slices"
)

func main() {
	common.Setup(20, part1, nil)
}

func part1(
	input string,
) string {
	p, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	dm, path, err := p.createDistanceMap()
	if err != nil {
		return fmt.Sprintf("Failed to create distance map: %v", err)
	}

	cheatCounts := make(map[int]int)
	for _, pos := range path {
		distance := dm[pos.Y][pos.X]
		next := []util.Coordinate{
			pos.North().North(),
			pos.South().South(),
			pos.East().East(),
			pos.West().West(),
		}

		for _, nPos := range next {
			// Normally moving two steps should give us two less distance,
			// so these 2 are subtracted from the saved amount.
			nDistance := dm[nPos.Y][nPos.X]
			saved := nDistance - distance - 2
			if saved > 0 {
				cheatCounts[saved] += 1
			}
		}
	}

	count := 0
	for k, v := range cheatCounts {
		if k >= 100 {
			count += v
		}
		fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
	}

	return fmt.Sprintf("Number of cheats saving at least 100ps: %d", count)
}

func (p parsedInput) createDistanceMap() (distanceMap, []util.Coordinate, error) {
	newM := make(distanceMap)
	path := make([]util.Coordinate, 0, p.length)
	for y, _ := range p.m {
		newM[y] = make(map[intX]int)
	}

	distance := 0
	curr := util.Coordinate{X: p.e.X, Y: p.e.Y}
	newM[curr.Y][curr.X] = 0

	visited := make(map[intY]map[intX]bool)
	for y := range p.m {
		visited[y] = make(map[intX]bool)
	}
	visited[curr.Y][curr.X] = true

	for !curr.Equals(p.s) {
		distance++
		found := false
		for _, newP := range curr.Adjacent4() {
			isBlocked := !p.m[newP.Y][newP.X]
			isVisited := visited[newP.Y][newP.X]
			if isBlocked || isVisited {
				continue
			}

			found = true
			newM[newP.Y][newP.X] = distance
			path = append(path, newP)
			visited[newP.Y][newP.X] = true
			curr = newP
			break
		}

		if !found {
			return newM, path, fmt.Errorf("missing path from %v", curr)
		}
	}

	// Skip end position and position next to end position,
	// because there's no point in cheating there.
	slices.Reverse(path[1:])

	return newM, path, nil
}
