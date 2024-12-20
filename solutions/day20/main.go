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
		for k, v := range findAllCheats(dm, pos, 2) {
			cheatCounts[k] += v
		}
	}

	count := 0
	for k, v := range cheatCounts {
		if k >= 100 {
			count += v
		}
	}

	return fmt.Sprintf("Number of cheats saving at least 100ps: %d", count)
}

func findAllCheats(
	dm distanceMap,
	pos util.Coordinate,
	steps int,
) map[int]int {
	// Initialize visited set
	visited := make(map[intY]map[intX]bool)
	for y := range dm {
		visited[y] = make(map[intX]bool)
	}

	cheats := make(map[int]int)
	positions := pos.Adjacent4()
	distance := dm[pos.Y][pos.X]
	for step := 1; step <= steps; step++ {
		newPositions := make([]util.Coordinate, 0, 4*len(positions))
		for _, p := range positions {
			// 1. Check if in visited set, if it is then skip.
			if visited[p.Y][p.X] || visited[p.Y] == nil {
				continue
			}
			visited[p.Y][p.X] = true

			// 2. Calculate time saved
			// Moving here normally will take $step number of steps, so savings
			// are calculated by subtracting step from the distance map.
			nDistance := dm[p.Y][p.X]
			saved := nDistance - distance - step

			if nDistance != 0 && saved <= 0 {
				// 3. If we didn't save time by going this way, continue.
				continue
			} else if nDistance != 0 {
				// 4. We didn't save time by going here, but we're still in the wilderness
				cheats[saved] += 1
			}

			// 5. This position is not trash, lets add all it's neighbors to the new positions set
			newPositions = append(newPositions, p.Adjacent4()...)
		}

		positions = newPositions
	}

	return cheats
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
