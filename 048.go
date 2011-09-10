package main

func Euler048() Result {
	acc := int64(0)
	for i := int64(1); i < 1001; i++ {
		acc += ModPow(i, i, 10000000000)
	}
	return acc % 10000000000
}
