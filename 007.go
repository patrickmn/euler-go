package main

import (
	"big"
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

func Euler007b() {
	var prime int
	ch := make(chan int)
	go Generate(ch)
	for num := 1; num < 10002; num++ {
		prime = <-ch
		ch1 := make(chan int)
		go FilterPrime(ch, ch1, prime)
		ch = ch1
	}
	result(prime)
}

// Faster
func Euler007() {
	numPrimes := 0
	for i := int64(2); ; i++ {
		if big.ProbablyPrime(big.NewInt(i), 4) {
			numPrimes += 1
			if numPrimes == 10001 {
				result(i)
				return
			}
		}
	}
}
