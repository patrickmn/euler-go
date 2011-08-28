package main

import (
	"strings"
)

var (
	words = map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
	hundred  = "hundred"
	thousand = "thousand"
)

func NumToWord(n int) string {
	if n == 1000 {
		return "one" + " " + thousand
	}
	hundreds := n / 100
	tens := (n - (hundreds * 100)) / 10
	ones := (n - (hundreds * 100)) - (tens * 10)

	str := ""
	if hundreds > 0 {
		str = words[hundreds] + " " + hundred
		if tens > 0 || ones > 0 {
			str = str + " and "
		}
	}
	if tens > 0 {
		if tens == 1 {
			str = str + words[(tens*10)+ones]
		} else {
			str = str + words[tens*10]
		}
	}
	if ones > 0 && tens != 1 {
		if tens > 0 {
			str = str + "-"
		}
		str = str + words[ones]
	}
	return str
}

func Euler017() {
	numLetters := 0
	for i := 1; i < 1001; i++ {
		word := NumToWord(i)
		word = strings.Replace(word, " ", "", -1)
		word = strings.Replace(word, "-", "", -1)
		numLetters += len(word)
	}
	result(numLetters)
}
