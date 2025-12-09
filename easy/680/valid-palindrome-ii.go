package leetcode

/*
Given a string s, return true if the s can be palindrome after deleting at most one character from it.
*/
func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindrome(s[:i]+s[i+1:]) || isPalindrome(s[:j]+s[j+1:])
		}
	}

	return true
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
