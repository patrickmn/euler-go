package main

func Euler003() {
	var (
		factors = []int64{}
		last    = int64(600851475143)
		c       = int64(2)
	)
	for last != 1 {
		if last%c == 0 && c%2 > 0 {
			factors = append(factors, c)
			last /= c
			c = c + 1
		} else {
			c = c + 1
		}
	}
	largest := int64(0)
	for _, v := range factors {
		if v > largest {
			largest = v
		}
	}
	result(largest)
}
