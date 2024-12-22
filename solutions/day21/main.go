package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

func main() {
	common.Setup(21, part1, nil)
}

func part1(
	input string,
) string {
	var out string

	sequences := strings.Split(input, "\n")
	sum := 0
	for _, code := range sequences {
		// Five is the longest path between two keys on numeric pad
		path := make([]int, 0, 5*len(code))
		prev := numStart
		for _, c := range code {
			curr := charToPosNum(int(c))
			path = append(path, shortestPath(prev, curr, true)...)
			prev = curr
		}

		for range 2 {
			// Three is the longest path between two keys on directional pad
			newPath := make([]int, 0, 3*len(path))
			prev = dirStart
			for _, c := range path {
				curr := charToPosDir(c)
				newPath = append(newPath, shortestPath(prev, curr, false)...)
				prev = curr
			}
			path = newPath
		}

		codeNumeric, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			return fmt.Sprintf("Failed to parse code %s as int: %v", code, err)
		}

		sum += len(path) * codeNumeric
	}

	out += fmt.Sprintf("Sum: %d", sum)
	return out
}
