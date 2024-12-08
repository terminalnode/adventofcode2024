package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(8, part1, part2)
}

func part1(
	input string,
) string {
	m, err := util.NewCharMatrix(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	coords := make(map[uint8][]util.Coordinate)
	for x := 0; x <= m.MaxX; x++ {
		for y := 0; y <= m.MaxY; y++ {
			char, _ := m.Get(x, y)
			if char == '.' {
				continue
			}
			coords[char] = append(coords[char], util.Coordinate{X: x, Y: y})
		}
	}

	antiNodeSet := make(map[string]bool)
	for _, charCoords := range coords {
		for i, thisCoord := range charCoords {
			for _, otherCoord := range charCoords[i+1:] {
				xDiff := thisCoord.X - otherCoord.X
				yDiff := thisCoord.Y - otherCoord.Y
				antiNodes := []util.Coordinate{
					{X: thisCoord.X + xDiff, Y: thisCoord.Y + yDiff},
					{X: otherCoord.X - xDiff, Y: otherCoord.Y - yDiff},
				}
				for _, antiNode := range antiNodes {
					if m.IsInMatrix(antiNode.X, antiNode.Y) {
						key := fmt.Sprintf("%d,%d", antiNode.X, antiNode.Y)
						antiNodeSet[key] = true
					}
				}
			}
		}
	}

	return fmt.Sprintf("Unique anti nodes: %d", len(antiNodeSet))
}

func part2(
	input string,
) string {
	m, err := util.NewCharMatrix(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	coords := make(map[uint8][]util.Coordinate)
	for x := 0; x <= m.MaxX; x++ {
		for y := 0; y <= m.MaxY; y++ {
			char, _ := m.Get(x, y)
			if char == '.' {
				continue
			}
			coords[char] = append(coords[char], util.Coordinate{X: x, Y: y})
		}
	}

	antiNodeSet := make(map[string]bool)
	for _, charCoords := range coords {
		for i, thisCoord := range charCoords {
			if len(charCoords) > 1 {
				addCoordinateToSet(antiNodeSet, thisCoord.X, thisCoord.Y)
			}

			for _, otherCoord := range charCoords[i+1:] {
				xDiff := thisCoord.X - otherCoord.X
				yDiff := thisCoord.Y - otherCoord.Y

				for n := 1; true; n++ {
					newX := thisCoord.X + (xDiff * n)
					newY := thisCoord.Y + (yDiff * n)
					if !m.IsInMatrix(newX, newY) {
						break
					}

					addCoordinateToSet(antiNodeSet, newX, newY)
				}

				for n := 1; true; n++ {
					newX := otherCoord.X - (xDiff * n)
					newY := otherCoord.Y - (yDiff * n)
					if !m.IsInMatrix(newX, newY) {
						break
					}

					addCoordinateToSet(antiNodeSet, newX, newY)
				}

			}
		}
	}

	return fmt.Sprintf("Unique anti nodes (with resonance): %d", len(antiNodeSet))
}

func addCoordinateToSet(
	set map[string]bool,
	x int,
	y int,
) {
	key := fmt.Sprintf("%02d,%02d", x, y)
	set[key] = true
}
