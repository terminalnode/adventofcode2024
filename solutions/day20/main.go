package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"slices"
)

func main() {
	common.Setup(20, part1, part2)
}

func part1(input string) string { return solve(input, 2, 100) }
func part2(input string) string { return solve(input, 20, 100) }

func solve(
	input string,
	steps int,
	limit int,
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
		for k, v := range findAllCheats(dm, pos, p.e, steps) {
			cheatCounts[k] += v
		}
	}

	count := 0
	for k, v := range cheatCounts {
		if k >= limit {
			count += v
		}
	}

	return fmt.Sprintf("Number of cheats saving at least %d ps: %d", limit, count)
}

func findAllCheats(
	dm distanceMap,
	pos util.Coordinate,
	ePos util.Coordinate,
	steps int,
) map[int]int {
	// Initialize visited set
	visited := make(map[intY]map[intX]bool)
	for y := range dm {
		visited[y] = make(map[intX]bool)
	}

	cheats := make(map[int]int)
	distance := dm[pos.Y][pos.X]
	for step := 2; step <= steps; step++ {
		for y := -step; y <= step; y++ {
			remainingSteps := step - util.AbsInt(y)
			for x := -remainingSteps; x <= remainingSteps; x++ {
				if util.AbsInt(x)+util.AbsInt(y) != step {
					continue
				}
				nPos := util.Coordinate{X: pos.X + x, Y: pos.Y + y}

				nDistance := dm[nPos.Y][nPos.X]
				saved := distance - nDistance - step
				if (nDistance != 0 || nPos.Equals(ePos)) && saved > 0 {
					cheats[saved] += 1
				}
			}
		}
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
