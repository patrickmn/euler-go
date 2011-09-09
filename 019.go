package main

var (
	MonthDays = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
)

func IsLeapYear(n int) bool {
	if n%100 == 0 {
		return n%400 == 0
	}
	return n%4 == 0
}

func Euler019() Result {
	// 1 Jan 1900 was a monday
	y := 1900
	m := 1
	d := 1
	wd := 1
	sundaysOnFirst := 0
	for !(y == 2000 && m == 12 && d == 31) {
		if wd == 7 {
			if d == 1 && y > 1900 {
				sundaysOnFirst++
			}
			wd = 1
		} else {
			wd++
		}
		if d >= MonthDays[m] {
			if m == 2 && IsLeapYear(y) && d == 28 {
				d++
			} else {
				d = 1
				if m == 12 {
					y++
					m = 1
				} else {
					m++
				}
			}
		} else {
			d++
		}
	}
	return sundaysOnFirst
}
