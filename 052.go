package main

import (
	"strconv"
)

func ContainsSameDigits(a, b int) bool {
	abytes := []byte{}
	astr := strconv.Itoa(a)
	for i := 0; i < len(astr); i++ {
		abytes = append(abytes, astr[i])
	}
	bbytes := []byte{}
	bstr := strconv.Itoa(b)
	for i := 0; i < len(bstr); i++ {
		bbytes = append(bbytes, bstr[i])
	}
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
