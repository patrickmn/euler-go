package main

import (
	"math/big"
)

func Euler055() Result {
	isLychrel := func(n int64) bool {
		cur := big.NewInt(n)
		for i := 0; i < 50; i++ {
			rev := new(big.Int)
			rev.SetString(ReverseString(cur.String()), 10)
			sum := new(big.Int)
			sum.Add(cur, rev)
			if IsPalindrome(sum.String()) {
				return false
			}
			cur.Set(sum)
		}
		return true
	}
	count := 0
	for i := int64(1); i < 10000; i++ {
		if isLychrel(i) {
			count++
		}
	}
	return count
}
