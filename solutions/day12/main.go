package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"log"
	"strings"
)

type plotMap = [][]plot
type plot struct {
	Label int
	Hash  string
}

func main() {
	common.Setup(12, part1, part2)
}

func part1(
	input string,
) string {
	sum := 0
	fields := 0
	plots := parsePlots(input)
	visited := make(map[string]bool)

	rows := len(plots)
	for y := 0; y < rows; y++ {
		columns := len(plots[y])
		for x := 0; x < columns; x++ {
			c := util.Coordinate{X: x, Y: y}
			area, perimeter, _ := countField(plots, c, visited)
			sum += area * perimeter

			if area != 0 {
				label := plots[y][x].Label
				log.Printf("Field %c at (%d,%d), area=%d, perim=%d", label, x, y, area, perimeter)
				fields++
			}
		}
	}

	return fmt.Sprintf("Area * perimeter of %d fields: %d", fields, sum)
}

func part2(
	input string,
) string {
	sum := 0
	fields := 0
	plots := parsePlots(input)
	visited := make(map[string]bool)

	rows := len(plots)
	for y := 0; y < rows; y++ {
		columns := len(plots[y])
		for x := 0; x < columns; x++ {
			c := util.Coordinate{X: x, Y: y}
			area, _, corners := countField(plots, c, visited)
			sum += area * corners

			if area != 0 {
				label := plots[y][x].Label
				log.Printf("Field %c at (%d,%d), area=%d, corners=%d", label, x, y, area, corners)
				fields++
			}
		}
	}

	return fmt.Sprintf("Area * sides of %d fields: %d", fields, sum)
}

func parsePlots(
	input string,
) plotMap {
	lines := strings.Split(input, "\n")
	m := make([][]plot, len(lines))

	for y, row := range lines {
		m[y] = make([]plot, len(row))
		for x, char := range row {
			c := util.Coordinate{X: x, Y: y}
			p := plot{Label: int(char), Hash: c.String()}
			m[y][x] = p
		}
	}

	return m
}

func countField(
	plots plotMap,
	c util.Coordinate,
	visited map[string]bool,
) (int, int, int) {
	if visited[c.String()] {
		return 0, 0, 0
	}
	visited[c.String()] = true
	label := plots[c.Y][c.X].Label

	area := 1
	perimeter := 0

	n := c.North()
	e := c.East()
	s := c.South()
	w := c.West()
	ne := c.NorthEast()
	nw := c.NorthWest()
	se := c.SouthEast()
	sw := c.SouthWest()

	nSame := isSameField(plots, label, n)
	eSame := isSameField(plots, label, e)
	sSame := isSameField(plots, label, s)
	wSame := isSameField(plots, label, w)
	neSame := isSameField(plots, label, ne)
	nwSame := isSameField(plots, label, nw)
	seSame := isSameField(plots, label, se)
	swSame := isSameField(plots, label, sw)

	// Haters gonna hate
	corners := countCorners(nSame, eSame, sSame, wSame, neSame, nwSame, seSame, swSame)

	if nSame {
		a, p, c := countField(plots, n, visited)
		area += a
		perimeter += p
		corners += c
	} else {
		perimeter++
	}

	if eSame {
		a, p, c := countField(plots, e, visited)
		area += a
		perimeter += p
		corners += c
	} else {
		perimeter++
	}

	if sSame {
		a, p, c := countField(plots, s, visited)
		area += a
		perimeter += p
		corners += c
	} else {
		perimeter++
	}

	if wSame {
		a, p, c := countField(plots, w, visited)
		area += a
		perimeter += p
		corners += c
	} else {
		perimeter++
	}

	return area, perimeter, corners
}

func isSameField(
	plots plotMap,
	label int,
	c util.Coordinate,
) bool {
	if !util.In2DArray(c, plots) {
		return false
	}
	cell := plots[c.Y][c.X]
	return cell.Label == label
}

func countCorners(
	nSame bool,
	eSame bool,
	sSame bool,
	wSame bool,
	neSame bool,
	nwSame bool,
	seSame bool,
	swSame bool,
) int {
	sum := 0
	if !nSame && !eSame {
		sum++
	} else if nSame && eSame && !neSame {
		sum++
	}

	if !nSame && !wSame {
		sum++
	} else if nSame && wSame && !nwSame {
		sum++
	}

	if !sSame && !eSame {
		sum++
	} else if sSame && eSame && !seSame {
		sum++
	}

	if !sSame && !wSame {
		sum++
	} else if sSame && wSame && !swSame {
		sum++
	}

	return sum
}
