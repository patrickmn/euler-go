package main

import (
	"strconv"
	"sync"
)

var (
	wg *sync.WaitGroup
)

func ReverseString(input string) string {
	n := 0
	rune := make([]int, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	return output
}

// Solution 1
func FilterPalindrome(ch chan<- int, a int) {
	var (
		product     int
		str, revstr string
	)
	for b := 999; b > 99; b-- {
		product = a * b
		str = strconv.Itoa(product)
		revstr = ReverseString(str)
		if str == revstr {
			ch <- product
		}
	}
	wg.Done()
}

func Euler004b() {
	var (
		largest int
		ch      = make(chan int, 100)
	)
	wg = &sync.WaitGroup{}
	for a := 999; a > 99; a-- {
		wg.Add(1)
		go FilterPalindrome(ch, a)
	}
	for {
		v := <-ch
		if v > largest {
			largest = v
		}
	}
	wg.Wait()
	result(largest)
}

// Solution 2
func Euler004() {
	var (
		largest     int
		product     int
		str, revstr string
	)
	for a := 999; a > 99; a-- {
		for b := 999; b > 99; b-- {
			product = a * b
			str = strconv.Itoa(product)
			revstr = ReverseString(str)
			if str == revstr {
				if product > largest {
					largest = product
				}
			}
		}
	}
	result(largest)
}
