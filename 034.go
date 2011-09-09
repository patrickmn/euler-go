package main

import (
	"strconv"
)

func Euler034() Result {
	equal := map[int]bool{}
	for i := 3; i < 50000; i++ {
		str := strconv.Itoa(i)
		fSum := 0
		for _, v := range str {
			num, _ := strconv.Atoi(string(v))
			fSum += Factorial(num)
		}
		if i == fSum {
			equal[i] = true
		}
	}
	sum := 0
	for k, _ := range equal {
		sum += k
	}
	return sum
}
