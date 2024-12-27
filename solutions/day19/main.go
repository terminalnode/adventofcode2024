package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"regexp"
	"strings"
)

func main() {
	common.Setup(19, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	p, err := parse(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	join := strings.Join(p.available, "|")
	r := regexp.MustCompile(fmt.Sprintf("^(%s)+$", join))

	count := 0
	for _, desired := range p.desired {
		if r.MatchString(desired) {
			count++
		}
	}

	return util.FormatAocSolution("%d of the %d desired designs are possible", count, len(p.desired))
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	p, err := parse(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	cache := make(map[string]int)
	count := 0
	for _, desired := range p.desired {
		n := loop(p.available, desired, cache)
		count += n
	}

	return util.FormatAocSolution("Number of available combinations: %d", count)
}

func loop(
	available []string,
	target string,
	cache map[string]int,
) int {
	cacheValue, isCached := cache[target]
	if isCached {
		return cacheValue
	}

	l := len(target)
	sum := 0
	for _, pattern := range available {
		pl := len(pattern)
		if pl > l {
			continue
		}

		if pattern == target {
			sum += 1
		} else if target[:pl] == pattern {
			sum += loop(available, target[pl:], cache)
		}
	}

	cache[target] = sum
	return sum
}
