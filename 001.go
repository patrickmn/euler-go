package main

func Euler001() {
	sum := 0
	for v := 1; v < 1000; v++ {
		if v%3 == 0 || v%5 == 0 {
			sum += v
		}
	}
	result(sum)
}
