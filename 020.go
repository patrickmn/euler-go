package main

import (
	"big"
	"strconv"
)

func Euler020() Result {
	z := big.NewInt(100)
	for i := 100; i > 0; i-- {
		z = z.Mul(z, big.NewInt(int64(i)))
	}
	sum := 0
	for _, v := range z.String() {
		num, _ := strconv.Atoi(string(v))
		sum += num
	}
	return sum
}
