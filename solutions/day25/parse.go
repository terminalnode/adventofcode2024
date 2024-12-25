package main

import (
	"strconv"
	"strings"
)

type entryType int
type keyArray []entry
type lockArray []entry

const (
	KEY = entryType(iota)
	LOCK
)

type entry struct {
	t    entryType
	a    []int
	aInv []int
	s    string
	sInv string
}

func parse(
	input string,
) (keyArray, lockArray) {
	split := strings.Split(input, "\n\n")
	keys := make(keyArray, 0, len(split))
	locks := make(lockArray, 0, len(split))

	for _, rawEntry := range split {
		lines := strings.Split(rawEntry, "\n")

		a := []int{0, 0, 0, 0, 0}
		aInv := []int{7, 7, 7, 7, 7}
		for _, line := range lines {
			for x, ch := range line {
				if ch == '#' {
					a[x] += 1
					aInv[x] -= 1
				}
			}
		}
		s := arrToString(a)
		sInv := arrToString(aInv)

		if lines[0] == "#####" {
			locks = append(locks, entry{
				t: LOCK,
				a: a, aInv: aInv,
				s: s, sInv: sInv})
		} else {
			keys = append(keys, entry{
				t: KEY,
				a: a, aInv: aInv,
				s: s, sInv: sInv})
		}
	}

	return keys, locks
}

func arrToString(
	arr []int,
) string {
	sArr := make([]string, len(arr))
	for i, n := range arr {
		sArr[i] = strconv.Itoa(n)
	}
	return strings.Join(sArr, ",")
}
