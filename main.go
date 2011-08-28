package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FuncType func()

var showAnswer func(k string) bool

func main() {
	funclist := map[string]FuncType{
		"1": Euler001,
		"2": Euler002,
		"3": Euler003,
		"4": Euler004,
		// "4b":  Euler004b,
		"5":   Euler005,
		"5b":  Euler005b,
		"6":   Euler006,
		"7":   Euler007,
		"7b":  Euler007b,
		"8":   Euler008,
		"9":   Euler009,
		"10":  Euler010,
		"10b": Euler010b,
		// "10c": Euler010c,
		"11": Euler011,
		"12": Euler012,
		"13": Euler013,
		"14": Euler014,
		"15": Euler015,
		"16": Euler016,
		"17": Euler017,
		"19": Euler019,
		"20": Euler020,
		"21": Euler021,
		"23": Euler023,
		"26": Euler026, // Not working
		"27": Euler027,
		"29": Euler029, // Not working
	}
	showAnswer = func(k string) bool {
		_, ok := funclist[k]
		if ok {
			fmt.Printf("Problem %3s: ", k)
			startTime := time.Nanoseconds()
			funclist[k]()
			endTime := time.Nanoseconds()
			elapsedms := (endTime - startTime) / 1e6
			extra := ""
			if elapsedms > 5000 {
				extra = "(!)"
			}
			fmt.Printf(" in %d ms%s\n", elapsedms, extra)
		}
		return ok
	}
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if arg == "all" {
			for i := 1; ; i++ {
				if !showAnswer(strconv.Itoa(i)) {
					break
				}
			}
		} else {
			_, ok := funclist[arg]
			if ok {
				showAnswer(arg)
			} else {
				fmt.Println("Invalid problem number")
			}
		}
	} else {
		name := os.Args[0]
		fmt.Println("Usage:", name, "<problem number, e.g. 1> or", name, "all")
	}
}
