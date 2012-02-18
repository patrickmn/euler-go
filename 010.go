package main

import (
	"github.com/jbarham/primegen.go"
	"math/big"
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

func Euler010b() Result {
	sum := int64(2) // 2 is a prime
	for i := int64(3); i < 2000000; i += 2 {
		if big.NewInt(i).ProbablyPrime(4) {
			sum += i
		}
	}
	return sum
}

func Euler010c() Result {
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

func Euler010d() Result {
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

func Euler010() Result {
	sieve := primegen.New()
	sum := uint64(0)
	for {
		prime := sieve.Next()
		if prime >= 2000000 {
			break
		}
		sum += prime
	}
	return sum
}
