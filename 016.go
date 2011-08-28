package main

func Euler016() {
	n := float64(2)
	for i := 1; i < 1000; i++ {
		n = n * 2
	}
	sum := SumOfIntsInFloat64(n)
	result(sum)
}
