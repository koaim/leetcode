package leetcode

/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/

func lengthOfLastWord(s string) int {
	var length int
	var waitLetter bool

	for _, v := range s {
		if string(v) == " " {
			waitLetter = true
		} else if waitLetter {
			length = 1
			waitLetter = false
		} else {
			length++
		}
	}

	return length
}
