package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(8, part1, nil)
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
	for c, charCoords := range coords {
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
