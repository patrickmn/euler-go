package main

import (
	"math/big"
)

// Faster
func Euler048() Result {
	acc := int64(0)
	for i := int64(1); i < 1001; i++ {
		acc += ModPow(i, i, 10000000000)
	}
	return acc % 10000000000
}

func Euler048b() Result {
	sum := new(big.Int)
	pow := new(big.Int)
	num := new(big.Int)
	for i := int64(1); i < 1001; i++ {
		num.SetInt64(i)
		pow.Exp(num, num, nil)
		sum.Add(sum, pow)
	}
	str := sum.String()
	return str[len(str)-10 : len(str)]
}
