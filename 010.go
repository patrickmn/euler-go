package main

import (
	"big"
)

var primes []int64

func isPrime(num int64) bool {
	for _, v := range primes {
		if num%v == 0 {
			return false
		}
	}
	return true
}

func Euler010() Result {
	num := int64(3)
	primes = append(primes, 2)
	for num < 2000000 {
		if big.ProbablyPrime(big.NewInt(num), 4) {
			primes = append(primes, num)
		}
		num += 2
	}
	return Sum64(primes)
}

func Euler010b() Result {
	num := int64(3)
	primes = append(primes, 2)
	for num < 2000000 {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num += 2
	}
	return Sum64(primes)
}

func Euler010c() Result {
	sum := 0
	ch := make(chan int)
	go Generate(ch)
	for num := 1; num < 2000000; num++ {
		prime := <-ch
		sum += prime
		ch1 := make(chan int)
		go FilterPrime(ch, ch1, prime)
		ch = ch1
	}
	return sum
}
