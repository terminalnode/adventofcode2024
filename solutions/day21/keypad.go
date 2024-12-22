package main

import "github.com/terminalnode/adventofcode2024/common/util"

type char = int32
type pathOptions = [][]char
type costMap map[char]map[char]pathOptions
type subCostMap map[char][][]char

var kpDirectional = buildKeypadCostMatrix([][]char{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
})

var kpNumeric = buildKeypadCostMatrix([][]char{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
})

func buildKeypadCostMatrix(
	keypad [][]char,
) costMap {
	out := make(costMap)
	for y1, row1 := range keypad {
		for x1, ch1 := range row1 {
			if ch1 == ' ' {
				continue
			}
			out[ch1] = make(subCostMap)

			for y2, row2 := range keypad {
				for x2, ch2 := range row2 {
					if ch2 == ' ' {
						continue
					} else if ch1 == ch2 {
						out[ch1][ch2] = [][]char{{ch1}}
						continue
					}
					start := util.Coordinate{X: x1, Y: y1}
					end := util.Coordinate{X: x2, Y: y2}
					out[ch1][ch2] = findPaths(keypad, start, end)
				}
			}
		}
	}
	return out
}

func findPaths(
	keypad [][]char,
	start util.Coordinate,
	end util.Coordinate,
) pathOptions {
	out := make(pathOptions, 0, 2)
	dy := start.Y - end.Y
	dx := start.X - end.X
	ady := util.AbsInt(dy)
	adx := util.AbsInt(dx)
	if ady > 0 && keypad[start.Y][end.X] != '' {
	}

	return out
}
