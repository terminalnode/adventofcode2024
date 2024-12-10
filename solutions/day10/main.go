package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strings"
)

func main() {
	common.Setup(10, part1, part2)
}

func part1(
	input string,
) string {
	topo := parseTopographicMap(input)

	out := 0
	for y := 0; y < len(topo); y++ {
		row := topo[y]
		for x := 0; x < len(row); x++ {
			value := topo[y][x]
			if value != 0 {
				continue
			}

			nineMap := make(map[string]bool)
			buildMapForPart1(topo, nineMap, util.Coordinate{X: x, Y: y}, -1)
			out += len(nineMap)
		}
	}

	return fmt.Sprintf("Number of trails: %d", out)
}

func part2(
	input string,
) string {
	topo := parseTopographicMap(input)

	out := 0
	for y := 0; y < len(topo); y++ {
		row := topo[y]
		for x := 0; x < len(row); x++ {
			value := topo[y][x]
			if value != 0 {
				continue
			}

			out += countDistinctTrails(topo, util.Coordinate{X: x, Y: y}, 0)
		}
	}

	return fmt.Sprintf("Number of distinct trails: %d", out)
}

func parseTopographicMap(
	input string,
) [][]int {
	lines := strings.Split(input, "\n")
	out := make([][]int, len(lines))
	for y, row := range lines {
		out[y] = make([]int, len(row))
		for x, char := range row {
			digit := int(char) - '0'
			out[y][x] = digit
		}
	}

	return out
}

func buildMapForPart1(
	topographicMap [][]int,
	nineMap map[string]bool,
	position util.Coordinate,
	prevValue int,
) {
	if !util.In2DArray(position, topographicMap) {
		return
	}

	value := topographicMap[position.Y][position.X]
	if value != prevValue+1 {
		return
	}

	if value == 9 {
		k := fmt.Sprintf("%v", position)
		nineMap[k] = true
	} else {
		cs := []util.Coordinate{
			position.North(),
			position.East(),
			position.South(),
			position.West(),
		}
		for _, newPos := range cs {
			buildMapForPart1(topographicMap, nineMap, newPos, value)
		}
	}
}

func countDistinctTrails(
	topographicMap [][]int,
	position util.Coordinate,
	value int,
) int {
	if value == 9 {
		return 1
	}

	cs := []util.Coordinate{
		position.North(),
		position.East(),
		position.South(),
		position.West(),
	}

	out := 0
	for _, c := range cs {
		if !util.In2DArray(c, topographicMap) {
			continue
		}

		cValue := topographicMap[c.Y][c.X]
		if cValue == value+1 {
			out += countDistinctTrails(topographicMap, c, cValue)
		}
	}

	return out
}
