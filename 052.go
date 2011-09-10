package main

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
