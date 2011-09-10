package main

import (
	"strconv"
)

func ContainsSameDigits(a, b int) bool {
	abytes := []byte(strconv.Itoa(a))
	bbytes := []byte(strconv.Itoa(b))
	for _, v := range abytes {
		found := false
		for _, ov := range bbytes {
			if v == ov {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Euler052() Result {
	isAnswer := func(n int) bool {
		for c := 2; c < 7; c++ {
			if !ContainsSameDigits(c*n, n) {
				return false
			}
		}
		return true
	}
	for i := 1; ; i++ {
		if isAnswer(i) {
			return i
		}
	}
	return 0
}
