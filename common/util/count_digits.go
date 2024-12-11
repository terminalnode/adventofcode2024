package util

import "strconv"

// This way is supposedly better than trying to be clever. Itoa is fast.
// https://stackoverflow.com/questions/68122675/fastest-way-to-find-number-of-digits-of-an-integer

func CountDigits(i int) int {
	return len(strconv.Itoa(i))
}
