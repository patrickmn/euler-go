package main

func Euler035() Result {
	numCP := 1 // 2 is a circular prime
	for i := 3; i < 1000000; i += 2 {
		if IsCircularPrime(i) {
			numCP++
		}
	}
	return numCP
}
