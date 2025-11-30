package leetcode

/*
Given two strings s and t, return true if they are both one edit distance apart, otherwise return false.
A string s is said to be one distance apart from a string t if you can:

	Insert exactly one character into s to get t.
	Delete exactly one character from s to get t.
	Replace exactly one character of s with a different character to get t.
*/
func isOneEditDistance(s string, t string) bool {
	if s == t || len(t)-len(s) > 1 {
		return false
	}

	if len(s) > len(t) {
		return isOneEditDistance(t, s)
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}

	return len(t)-len(s) == 1
}
