package leetcode

/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.
*/

import "unicode/utf8"

func findAnagrams(s string, p string) []int {
	m := map[rune]int{}

	for _, v := range p {
		m[v]++
	}

	k := utf8.RuneCountInString(p)
	r := []rune(s)
	res := make([]int, 0)

	for i := 0; i < len(r)-k+1; i++ {
		mm := map[rune]int{}
		for j := i; j < i+k; j++ {
			if _, ok := m[r[j]]; ok {
				mm[r[j]]++
			}
		}

		if len(mm) == len(m) {
			b := true
			for k, v := range mm {
				if v != m[k] {
					b = false
				}
			}
			if b {
				res = append(res, i)
			}

		}
	}

	return res
}
