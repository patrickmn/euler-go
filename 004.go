package main

import (
	"strconv"
	"sync"
)

var (
	wg *sync.WaitGroup
)

// Solution 1
func FilterPalindrome(ch chan<- int, a int) {
	var product     int
	for b := 999; b > 99; b-- {
		product = a * b
		if IsPalindrome(strconv.Itoa(product)) {
			ch <- product
		}
	}
	wg.Done()
}

func Euler004b() Result {
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
	return largest
}

// Solution 2
func Euler004() Result {
	var (
		largest     int
		product     int
	)
	for a := 999; a > 99; a-- {
		for b := 999; b > 99; b-- {
			product = a * b
			if IsPalindrome(strconv.Itoa(product)) {
				if product > largest {
					largest = product
				}
			}
		}
	}
	return largest
}
