package main

import (
	"strconv"
)

func Euler036() Result {
	sum := 0
	for i := 0; i < 1000000; i++ {
		// Palindromic in base 10 (decimal)?
		if IsPalindrome(strconv.Itoa(i)) {
			binarystr := strconv.Itob(i, 2)
			if binarystr[:1] == "0" {
				continue
			} else if IsPalindrome(binarystr) {
				sum += i
			}
		}
	}
	return sum
}
