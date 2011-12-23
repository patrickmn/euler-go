package main

import (
	"strconv"
)

func GridRoutes(n, m float64) float64 {
	return FactorialF64(m+n) / (FactorialF64(m) * FactorialF64(n))
}

func Euler015() Result {
	possible := GridRoutes(20, 20)
	return strconv.FormatFloat(possible, 'f', 0, 64)
}
