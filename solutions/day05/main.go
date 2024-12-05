package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

type ruleSet = map[int][]int
type pageList = []int

type parsedInput = struct {
	rules   ruleSet
	manuals []pageList
}

func main() {
	common.Setup(5, part1, nil)
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

func divideCorrectManuals(
	input parsedInput,
) ([]pageList, []pageList) {
	var correct []pageList
	var incorrect []pageList

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
	rules ruleSet,
	pages pageList,
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
