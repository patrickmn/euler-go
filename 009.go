package main

func checkTriplet(sum, a, b, c int) bool {
	return (a*a)+(b*b) == (c*c) && a+b+c == sum
}

func findTriplet(sum int) int {
	var a, b, c int
	for a = 0; a < 1000; a++ {
		for b = a + 1; b < 1000; b++ {
			for c = b + 1; c < 1000; c++ {
				if checkTriplet(sum, a, b, c) {
					return a * b * c
				}
			}
			if checkTriplet(sum, a, b, c) {
				return a * b * c
			}
		}
		if checkTriplet(sum, a, b, c) {
			return a * b * c
		}
	}
	return -1
}

func Euler009() {
	result(findTriplet(1000))
}
