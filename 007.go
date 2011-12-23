package main

import (
	"math/big"
	primegen "github.com/jbarham/primegen.go"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func FilterPrime(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Euler007c() Result {
	var prime int
	ch := make(chan int)
	go Generate(ch)
	for num := 1; num < 10002; num++ {
		prime = <-ch
		ch1 := make(chan int)
		go FilterPrime(ch, ch1, prime)
		ch = ch1
	}
	return prime
}

// Faster
func Euler007b() Result {
	numPrimes := 0
	for i := int64(2); ; i++ {
		if big.ProbablyPrime(big.NewInt(i), 2) {
			numPrimes += 1
			if numPrimes == 10001 {
				return i
			}
		}
	}
	return 0
}

// Even faster
func Euler007() Result {
	sieve := primegen.New()
	for i := 0; i < 10000; i++ {
		sieve.Next()
	}
	return sieve.Next()
}
