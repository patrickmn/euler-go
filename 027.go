package main

import (
	"math/big"
)

func Euler027() Result {
	var maxConsecutive, maxConsecutiveA, maxConsecutiveB int64
	for a := int64(-1000); a < 0; a++ {
		for b := int64(0); b < 1000; b++ {
			consecutive := int64(0)
			for n := int64(0); n < 80; n++ {
				num := (n * n) + (a * n) + b
				if big.NewInt(num).ProbablyPrime(1) {
					consecutive++
				}
			}
			if consecutive > maxConsecutive {
				maxConsecutive = consecutive
				maxConsecutiveA = a
				maxConsecutiveB = b
			}
		}
	}
	product := maxConsecutiveA * maxConsecutiveB
	return product
}
