package main

func Euler021() {
	pairs := map[int]int{}
	for i := 1; i < 10000; i++ {
		dsum := Sum(Divisors(i))
		odsum := Sum(Divisors(dsum))
		if i == odsum && odsum != dsum {
			_, found := pairs[dsum]
			if !found {
				pairs[i] = dsum
			}
		}
	}
	sum := 0
	for k, v := range pairs {
		sum += k + v
	}
	result(sum)
}
