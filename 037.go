package main

func Euler037() Result {
	sum := 0
	count := 0
	// 2, 3, 5, and 7 are not considered truncatable primes
	for i := 11; ; i += 2 {
		if IsTruncatablePrime(i) {
			count++
			sum += i
			if count == 11 {
				break
			}
		}
	}
	return sum
}
