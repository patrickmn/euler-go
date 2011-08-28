package main

import (
	"fmt"
	"math"
)

func Euler029() {
	nums := map[float64]bool{}
	for a := float64(2); a <= 100; a++ {
		for b := float64(2); b <= 100; b++ {
			num := math.Pow(a, b)
			nums[num] = true
			fmt.Println(a, b)
		}
	}
	fmt.Println(len(nums))
}
