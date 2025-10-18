package leetcode

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/
func isPalindrome(s string) bool {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	str := nonAlphanumericRegex.ReplaceAllString(strings.ReplaceAll(strings.ToLower(s), " ", ""), "")

	len := utf8.RuneCountInString(str)
	runes := []rune(str)

	result := true
	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			result = false
			break
		}
	}

	return result
}
