package main

import (
	"fmt"
	"math"
)

func Euler029() Result {
	nums := map[float64]bool{}
	for a := float64(2); a <= 100; a++ {
		for b := float64(2); b <= 100; b++ {
			num := math.Pow(a, b)
			nums[num] = true
			fmt.Println(a, b)
		}
	}
	return len(nums)
}
