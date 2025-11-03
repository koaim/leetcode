package leetcode

import "unicode/utf8"

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if utf8.RuneCountInString(s) > utf8.RuneCountInString(t) {
		return false
	}

	sR := []rune(s)
	tR := []rune(t)

	var sPtr, tPtr, match int
	for {
		if match == utf8.RuneCountInString(s) || tPtr == len(tR) {
			break
		}
		if sR[sPtr] == tR[tPtr] {
			match++
			sPtr++
		}
		tPtr++
	}

	return match == utf8.RuneCountInString(s)
}
