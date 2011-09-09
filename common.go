package main

import (
	"fmt"
	"sort"
	"strconv"
)

func GetString(a interface{}) string {
        return fmt.Sprintf("%v", a)
}

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

func IsDeficient(n int) bool {
	return Sum(Divisors(n)) < n
}

func IsPerfect(n int) bool {
	return Sum(Divisors(n)) == n
}

func IsAbundant(n int) bool {
	return Sum(Divisors(n)) > n
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
