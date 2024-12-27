package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strconv"
	"strings"
)

func main() { common.Setup(21, part1, part2) }

func part1(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, 2) }
func part2(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, 25) }

func solve(
	input util.AocInput,
	robots int,
) (util.AocSolution, util.AocError) {
	cache := make(cacheMap)

	sequences := strings.Split(input.Input, "\n")
	sum := 0
	for _, code := range sequences {
		codeNumeric, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			msg := fmt.Sprintf("Failed to parse code %s as int: %v", code, err)
			return util.NewAocError(msg, util.StringToNumber)
		}

		// Five is the longest path between two keys on numeric pad
		path := make([][]byte, 0, 5*len(code))
		prev := numStart
		for _, c := range code {
			curr := charToPosNum(byte(c))
			path = append(path, shortestPath(prev, curr, true))
			prev = curr
		}

		pathLen := 0
		for _, c := range path {
			pathLen += dfs(cache, cacheKey{path: string(c), depth: robots})
		}

		sum += pathLen * codeNumeric
	}

	return util.FormatAocSolution("Sum: %d", sum)
}
