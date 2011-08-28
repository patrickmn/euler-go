package main

func Euler002() {
	var (
		a, b = 0, 1
		sum  = 0
	)
	for b <= 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	result(sum)
}
