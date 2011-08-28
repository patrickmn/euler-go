package main

func Euler006() {
	var (
		sumNumbers = 0
		sumSquares = 0
		difference int
	)
	for num := 1; num < 101; num++ {
		sumNumbers += num
		sumSquares += num * num
	}
	difference = (sumNumbers * sumNumbers) - sumSquares
	result(difference)
}
