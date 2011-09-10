package main

import (
	"big"
	"strconv"
)

func IsCircularPrime(n int) bool {
	ints := []byte{}
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		ints = append(ints, str[i])
	}
	top := len(ints)
	for oi := 0; oi < top; oi++ {
		prev := ints[0]
		for i := 0; i < top; i++ {
			if i+1 < top {
				prev, ints[i+1] = ints[i+1], prev
			} else {
				ints[0] = prev
			}
		}
		num, _ := strconv.Atoi(string(ints))
		if !big.ProbablyPrime(big.NewInt(int64(num)), 1) {
			return false
		}
	}
	return true
}

func Euler035() Result {
	numCP := 0
	for n := 1; n < 1000000; n++ {
		if IsCircularPrime(n) {
			numCP++
		}
	}
	return numCP
}
