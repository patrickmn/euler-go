package main

import (
	"strconv"
)

func Euler036() Result {
	sum := 0
	for i := 0; i < 1000000; i++ {
		// We can skip even numbers because their binary representation always ends in 0
		// (even `mod` 1 == 0), thus the reversed base 2 would have a leading zero.
		if i%2 == 0 {
			continue
		}
		// Palindromic in base 10 (decimal)?
		if IsPalindrome(strconv.Itoa(i)) {
			binarystr := strconv.Itob(i, 2)
			// Leading zero not allowed
			if binarystr[:1] == "0" {
				continue
			} else if IsPalindrome(binarystr) {
				sum += i
			}
		}
	}
	return sum
}
