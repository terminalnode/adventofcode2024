package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"regexp"
	"strings"
)

func main() {
	common.Setup(19, part1, part2)
}

func part1(
	input string,
) string {
	p, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	join := strings.Join(p.available, "|")
	r := regexp.MustCompile(fmt.Sprintf("^(%s)+$", join))

	count := 0
	for _, desired := range p.desired {
		if r.MatchString(desired) {
			count++
		}
	}

	return fmt.Sprintf("%d of the %d desired designs are possible", count, len(p.desired))
}

func part2(
	input string,
) string {
	p, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	cache := make(map[string]int)
	count := 0
	for _, desired := range p.desired {
		n := loop(p.available, desired, cache)
		count += n
	}

	return fmt.Sprintf("Number of available combinations: %d", count)
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
