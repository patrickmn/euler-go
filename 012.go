package main

func Over500Divisors(n int64) bool {
	// The number itself is always a factor
	numFactors := 1
	for i := int64(1); i < int64(1000); i++ {
		if n%i == 0 {
			numFactors++
		}
	}
	if numFactors < 100 {
		// This one probably doesn't have 500 divisors; give up
		return false
	}
	for i := int64(1000); i < n; i++ {
		if n%i == 0 {
			numFactors++
		}
	}
	return numFactors > 500
}

func Euler012() Result {
	num := int64(0)
	for i := int64(1); ; i++ {
		num += i
		if Over500Divisors(num) {
			return num
		}
	}
	return 0
}
