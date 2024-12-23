package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

func main() {
	common.Setup(22, part1, part2)
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

type priceDiff struct {
	diff  int
	price int
}

func part2(
	input string,
) string {
	secrets, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	allKeys := make(map[string]bool)
	allPriceMaps := make([]map[string]int, len(secrets))

	for i, secret := range secrets {
		ds := make([]priceDiff, 2000)
		price := priceOf(secret)
		visited := make(map[string]bool)
		priceMap := make(map[string]int)

		for n := 0; n < 2000; n++ {
			secret = evolveSecretNumber(secret)
			newPrice := priceOf(secret)
			ds[n] = priceDiff{diff: newPrice - price, price: newPrice}
			price = newPrice
		}

		for n := 3; n <= 1999; n++ {
			key := fmt.Sprintf(
				"%d,%d,%d,%d",
				ds[n-3].diff, ds[n-2].diff, ds[n-1].diff, ds[n].diff)
			if visited[key] {
				continue
			}

			allKeys[key] = true
			visited[key] = true
			priceMap[key] = ds[n].price
		}
		allPriceMaps[i] = priceMap
	}

	bestKey := ""
	bestSum := 0
	for key, _ := range allKeys {
		sum := 0
		for _, pm := range allPriceMaps {
			sum += pm[key]
		}

		if sum > bestSum {
			bestSum = sum
			bestKey = key
		}
	}

	return fmt.Sprintf("Max number of bananas %d (%s)", bestSum, bestKey)
}

func priceOf(n int) int {
	s := strconv.Itoa(n)
	ch := s[len(s)-1]
	return int(ch - '0')
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
