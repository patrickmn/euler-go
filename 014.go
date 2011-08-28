package main

func GetNumJumps(num int) int {
	jumps := 0
	n := int64(num)
	for {
		if n%2 == 0 {
			// Even
			n = n / 2
		} else {
			// Odd
			n = (3 * n) + 1
		}
		jumps++
		if n == 1 {
			break
		}
	}
	return jumps
}

func Euler014() {
	var (
		longestNum   = 0
		longestChain = 0
	)
	for i := 2; i < 1000000; i++ {
		jumps := GetNumJumps(i)
		if jumps > longestChain {
			longestChain = jumps
			longestNum = i
		}
	}
	result(longestNum)
}
