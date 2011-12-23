package main

import (
	"fmt"
	"strconv"
)

func Euler026() Result {
	var (
		strings = map[string]float64{}
		one     = float64(1)
	)
	for i := float64(1); i < float64(1000); i++ {
		frac := one / i
		str := strconv.FormatFloat(frac, 'f', 100, 64)
		strings[str] = i
	}
	//fmt.Println(strings)
	var (
		longestRecurring       int = 0
		longestRecurringString string
	)
	for str, _ := range strings {
		recurring := 0
		last := ""
		for _, v := range str {
			chr := string(v)
			// Need to find a _cycle_, not just recurrance of a single number!
			if last != "" && last == chr {
				recurring += 1
				if recurring > longestRecurring {
					longestRecurring = recurring
					longestRecurringString = str
				}
			} else {
				recurring = 0
			}
			last = chr
		}
		fmt.Println("str:", str, "recurring:", recurring)
	}
	return strings[longestRecurringString]
}
