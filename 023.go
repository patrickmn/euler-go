package main

func IsSumOfTwo(n int, a *[]int) bool {
	for _, v1 := range *a {
		for _, v2 := range *a {
			if v1+v2 == n {
				return true
			}
		}
	}
	return false
}

func GetAbundantNumbers() *[]int {
	// By mathematical analysis, it can be shown that all integers greater
	// than 28123 can be written as the sum of two abundant numbers. 
	var (
		abundant = []int{}
	)
	for i := 1; i < 28124; i++ {
		if IsAbundant(i) {
			abundant = append(abundant, i)
		}
	}
	return &abundant
}

func Euler023() Result {
	var (
		sum      = 0
		abundant = GetAbundantNumbers()
	)
	for i := 1; i < 28124; i++ {
		if !IsSumOfTwo(i, abundant) {
			sum += i
		}
	}
	return sum
}
