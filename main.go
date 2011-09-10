package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Result interface{}

type EulerFunc func() Result

var showAnswer func(k string)

func main() {
	// Functions that don't work yet are commented out
	solutionFuncs := map[string]EulerFunc{
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
		// "11": Euler011,
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
		// "26": Euler026,
		"27": Euler027,
		// "29": Euler029,
		"30":  Euler030,
		"34":  Euler034,
		"35":  Euler035,
		"36":  Euler036,
		"37":  Euler037,
		"40":  Euler040,
		"48":  Euler048,
		"48b": Euler048b,
		"52":  Euler052,
		"55":  Euler055,
	}
	answers := map[string]string{
		"1":  "233168",
		"2":  "4613732",
		"3":  "6857",
		"4":  "906609",
		"5":  "232792560",
		"6":  "25164150",
		"7":  "104743",
		"8":  "40824",
		"9":  "31875000",
		"10": "142913828922",
		// "11":
		"12": "76576500",
		"13": "5537376230",
		"14": "837799",
		"15": "137846528820",
		"16": "1366",
		"17": "21124",
		// "18":
		"19": "171",
		"20": "648",
		"21": "31626",
		"22": "871198282",
		"23": "4179871",
		// "24":
		"25": "4782",
		// "26":
		"27": "-59231",
		// "28":
		"29": "9183",
		"30": "443839",
		"34": "40730",
		"35": "55",
		"36": "872187",
		"37": "748317",
		"40": "210",
		"48": "9110846700",
		"52": "142857",
		"55": "249",
	}
	showAnswer = func(k string) {
		_, ok := solutionFuncs[k]
		if ok {
			fmt.Printf("Problem %3s: ", k)
			startTime := time.Nanoseconds()
			res := solutionFuncs[k]()
			endTime := time.Nanoseconds()
			elapsedms := (endTime - startTime) / 1e6
			error := ""
			realRes, ok := answers[k]
			if ok {
				if GetString(res) != realRes {
					error = "FAIL"
				}
			} else {
				error = "???"
			}
			extra := ""
			if elapsedms > 5000 {
				extra = "(!)"
			}
			fmt.Printf("%4s %12v in %d ms%s", error, res, elapsedms, extra)
			fmt.Println("")
		}
	}
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if arg == "all" {
			// Some problems might not be solved yet, so just naively try every
			// number below the top Euler problem.
			for i := 1; i < 500; i++ {
				showAnswer(strconv.Itoa(i))
			}
		} else {
			_, ok := solutionFuncs[arg]
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
