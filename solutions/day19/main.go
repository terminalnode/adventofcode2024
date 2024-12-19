package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"regexp"
	"strings"
)

func main() {
	common.Setup(19, part1, nil)
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
