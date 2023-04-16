package leetcode

import "unicode/utf8"

/*
Given a roman numeral, convert it to an integer.
*/
func romanToInt(s string) int {
	conv := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	exclusions := map[string]map[string]int{
		"I": {
			"V": 4,
			"X": 9,
		},
		"X": {
			"L": 40,
			"C": 90,
		},
		"C": {
			"D": 400,
			"M": 900,
		},
	}

	sum := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		v := string(runes[i])

		if utf8.RuneCountInString(s)-i > 1 {
			if excl, ok := exclusions[v]; ok {
				if val, ok2 := excl[string(runes[i+1])]; ok2 {
					sum += val
					i++
				} else {
					sum += conv[v]
				}
			} else {
				sum += conv[v]
			}
		} else {
			sum += conv[v]
		}
	}

	return sum
}
