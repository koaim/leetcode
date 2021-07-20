package leetcode

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
*/

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}

	return x == rev || x == rev/10
}
