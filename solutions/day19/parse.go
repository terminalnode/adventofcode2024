package main

import (
	"fmt"
	"strings"
)

type parsedInput struct {
	available []string
	desired   []string
}

func parse(
	input string,
) (parsedInput, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return parsedInput{}, fmt.Errorf("expected two parts in input, got %d", len(split))
	}

	available := strings.Split(split[0], ", ")
	if len(available) == 0 {
		return parsedInput{}, fmt.Errorf("no available patterns found")
	}

	desired := strings.Split(split[1], "\n")
	if len(desired) == 0 {
		return parsedInput{}, fmt.Errorf("no desired patterns found")
	}

	return parsedInput{
		available: available,
		desired:   desired,
	}, nil
}
