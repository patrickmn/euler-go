package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func Euler040() Result {
	buf := bytes.NewBufferString("")
	for i := 1; ; i++ {
		fmt.Fprint(buf, i)
		if buf.Len() > 1000000 {
			break
		}
	}
	str := buf.String()
	get := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	sum := 1
	for _, v := range get {
		num, _ := strconv.Atoi(str[v-1 : v])
		sum *= num
	}
	return sum
}
