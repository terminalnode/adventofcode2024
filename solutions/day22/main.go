package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

func main() {
	common.Setup(22, part1, nil)
}

func part1(
	input string,
) string {
	secrets, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := 0
	for _, secret := range secrets {
		for n := 0; n < 2000; n++ {
			secret = evolveSecretNumber(secret)
		}
		sum += secret
	}

	return fmt.Sprintf("Sum of all secret numbers after 2k rounds: %d", sum)
}

func parse(
	input string,
) ([]int, error) {
	lines := strings.Split(input, "\n")
	numbers := make([]int, 0, len(lines))
	for _, l := range lines {
		n, err := strconv.ParseInt(l, 10, 0)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, int(n))
	}
	return numbers, nil
}

func evolveSecretNumber(
	secretNumber int,
) int {
	secretNumber = pruneNumber((secretNumber * 64) ^ secretNumber)
	secretNumber = pruneNumber((secretNumber / 32) ^ secretNumber)
	return pruneNumber((secretNumber * 2048) ^ secretNumber)
}

func pruneNumber(n int) int {
	return n % 16777216
}
