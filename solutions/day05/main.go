package main

import (
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

type RuleSet = map[int][]int
type PageList = []int

type parsedInput = struct {
	rules   RuleSet
	manuals []PageList
}

func main() {
	common.Setup(5, part1, part2)
}

func part1(
	input string,
) string {
	parsed, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := 0
	filteredManuals, _ := divideCorrectManuals(parsed)
	for _, manual := range filteredManuals {
		sum += manual[len(manual)/2]
	}

	return fmt.Sprintf("Sum of middle numbers: %d", sum)
}

func part2(
	input string,
) string {
	parsed, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := 0
	_, filteredManuals := divideCorrectManuals(parsed)
	for _, manual := range filteredManuals {
		// Technically we don't need to correct the whole manual
		corrected, err := correctManual(parsed.rules, manual, PageList{})
		if err != nil {
			return fmt.Sprintf("Failed to correct manual %v:\n%v", manual, err)
		}
		sum += corrected[len(corrected)/2]
	}

	return fmt.Sprintf("Sum of middle numbers (corrected manuals): %d", sum)
}

func divideCorrectManuals(
	input parsedInput,
) ([]PageList, []PageList) {
	var correct []PageList
	var incorrect []PageList

	for _, manual := range input.manuals {
		if verifyOrder(input.rules, manual) {
			correct = append(correct, manual)
		} else {
			incorrect = append(incorrect, manual)
		}
	}

	return correct, incorrect
}

func verifyOrder(
	rules RuleSet,
	pages PageList,
) bool {
	visited := map[int]bool{}

	for _, page := range pages {
		for _, forbidden := range rules[page] {
			if visited[forbidden] {
				return false
			}
		}

		visited[page] = true
	}

	return true
}

func parseInput(
	input string,
) (parsedInput, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return parsedInput{}, fmt.Errorf("input should have only two parts, had %d", len(parts))
	}

	rules, err := parseRules(parts[0])
	if err != nil {
		return parsedInput{}, err
	}

	manuals, err := parsedManuals(parts[1])
	if err != nil {
		return parsedInput{}, err
	}

	return parsedInput{rules, manuals}, nil
}

func parseRules(
	input string,
) (map[int][]int, error) {
	lines := strings.Split(input, "\n")
	rules := map[int][]int{}

	for i, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return rules, fmt.Errorf("expected only two parts in rule #%d (%v)", i+1, parts)
		}
		k, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			return rules, fmt.Errorf("can't parse '%v' to int", parts[0])
		}
		key := int(k)

		v, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			return rules, fmt.Errorf("can't parse '%v' to int", parts[1])
		}
		value := int(v)

		rules[key] = append(rules[key], value)
	}

	return rules, nil
}

func parsedManuals(
	input string,
) ([][]int, error) {
	lines := strings.Split(input, "\n")
	manuals := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ",")
		pages := make([]int, len(parts))
		for i, pageNumber := range parts {
			v, err := strconv.ParseInt(pageNumber, 10, 0)
			if err != nil {
				return manuals, fmt.Errorf("can't parse '%v' to int", pageNumber)
			}
			pages[i] = int(v)
		}
		manuals[i] = pages
	}

	return manuals, nil
}

func correctManual(
	rules RuleSet,
	pageList PageList,
	result PageList,
) (PageList, error) {
	if len(pageList) == 0 {
		return result, nil
	}

	remaining := pageList
	for len(remaining) > 0 {
		for i, page := range remaining {
			if !pageIsAnOption(rules, remaining, page) {
				continue
			}

			rem := append(remaining[:i], remaining[i+1:]...)
			res := append(result, page)
			out, err := correctManual(rules, rem, res)
			if err == nil {
				return out, err
			}
		}
	}

	return PageList{}, errors.New("failure")
}

func pageIsAnOption(
	rules RuleSet,
	remaining PageList,
	this int,
) bool {
	for _, other := range remaining {
		if this == other {
			continue
		}

		for _, forbidden := range rules[other] {
			if this == forbidden {
				return false
			}
		}
	}

	return true
}
