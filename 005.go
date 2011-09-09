package main

// Solution 1
var (
	sem   = make(chan bool, 100)
	final int
)

func tryDivision(num int) {
	// Don't need to check for 20
	for i := 1; i < 20; i++ {
		if num%i != 0 {
			<-sem
			return
		}
	}
	final = num
}

func Euler005b() Result {
	for num := 20; ; num += 20 {
		if final != 0 {
			return final
		}
		sem <- true
		go tryDivision(num)
	}
	return 0
}

// Solution 2 - Faster
func canDivide(num int) bool {
	for i := 1; i < 20; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}

func Euler005() Result {
	for num := 20; ; num += 20 {
		if canDivide(num) {
			return num
		}
	}
	return 0
}
