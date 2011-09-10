package main

import (
	"big"
)

func Euler048() Result {
	acc := int64(0)
	for i := int64(1); i < 1001; i++ {
		acc += ModPow(i, i, 10000000000)
	}
	return acc % 10000000000
}

func Euler048b() Result {
	sum := big.NewInt(0)
	pow := big.NewInt(0)
	var num *big.Int
	for i := 1; i < 1001; i++ {
		num = big.NewInt(int64(i))
		pow.Exp(num, num, nil)
		sum.Add(sum, pow)
	}
	str := sum.String()
	return str[len(str)-10 : len(str)]
}