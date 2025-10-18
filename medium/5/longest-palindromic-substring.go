package leetcode

/*
Given a string s, return the longest palindromic substring in s.
*/

import "unicode/utf8"

func longestPalindrome(s string) string {
	if utf8.RuneCountInString(s) == 1 {
		return s
	}

	palindrom := []rune{}

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if len(palindrom) > len(r)-i {
			break
		}

		for j := i; j < len(r); j++ {
			s := r[i : j+1]
			if isPalindrom(s) && len(s) > len(palindrom) {
				palindrom = s
			}
		}
	}

	return string(palindrom)
}

func isPalindrom(r []rune) bool {
	if len(r) == 0 {
		return false
	}

	res := true
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			res = false
			break
		}
	}

	return res
}
