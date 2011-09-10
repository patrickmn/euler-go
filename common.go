package main

import (
	"big"
	"fmt"
	"sort"
	"strconv"
)

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return n * -1
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Sum64(a []int64) int64 {
	sum := int64(0)
	for _, v := range a {
		sum += v
	}
	return sum
}

func ModPow(base, ex, modulo int64) int64 {
	acc := int64(1)
	ex--
	for ex >= 0 {
		acc *= base
		acc %= modulo
		if acc <= 0 {
			return 0
		}
		ex--
	}
	return acc
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Factorial64(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * Factorial64(n-1)
}

func FactorialF64(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * FactorialF64(n-1)
}

func Divisors(n int) []int {
	divisors := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func SumOfIntsInInt(n int) int {
	str := strconv.Itoa(n)
	sum := 0
	for _, v := range str {
		num, _ := strconv.Atoi(string(v))
		sum += num
	}
	return sum
}


func SumOfIntsInFloat64(n float64) int {
	str := strconv.Ftoa64(n, 'f', 0)
	sum := 0
	for _, v := range str {
		num, _ := strconv.Atoi(string(v))
		sum += num
	}
	return sum
}

func ContainsSameDigits(a, b int) bool {
	abytes := []byte(strconv.Itoa(a))
	bbytes := []byte(strconv.Itoa(b))
	for _, v := range abytes {
		found := false
		for _, ov := range bbytes {
			if v == ov {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func IsDeficient(n int) bool {
	return Sum(Divisors(n)) < n
}

func IsPerfect(n int) bool {
	return Sum(Divisors(n)) == n
}

func IsAbundant(n int) bool {
	return Sum(Divisors(n)) > n
}

func IsCircularPrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}
	ints := []byte(strconv.Itoa(n))
	top := len(ints)
	for oi := 0; oi < top; oi++ {
		prev := ints[0]
		for i := 0; i < top; i++ {
			if i+1 < top {
				prev, ints[i+1] = ints[i+1], prev
			} else {
				ints[0] = prev
			}
		}
		num, _ := strconv.Atoi(string(ints))
		if !big.ProbablyPrime(big.NewInt(int64(num)), 1) {
			return false
		}
	}
	return true
}

func IsTruncatablePrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i:])
		if !big.ProbablyPrime(big.NewInt(int64(num)), 1) {
			return false
		}
	}
	return true
}

func IsDoubleTruncatablePrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		left, _ := strconv.Atoi(str[i:])
		right, _ := strconv.Atoi(str[:len(str)-i])
		if !big.ProbablyPrime(big.NewInt(int64(left)), 1) || !big.ProbablyPrime(big.NewInt(int64(right)), 1) {
			return false
		}
	}
	return true
}

func SortedKeys(items map[string]interface{}) []string {
	keys := make([]string, len(items))
	i := 0
	for key, _ := range items {
		keys[i] = key
		i++
	}
	sort.SortStrings(keys)
	return keys
}

func GetString(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

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

func IsPalindrome(s string) bool {
	rev := ReverseString(s)
	return s == rev
}
