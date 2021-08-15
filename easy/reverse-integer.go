package leetcode

/*
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range, then return 0.
*/

import "math"

func reverse(x int) int {
	var rev int

	for x != 0 {
		p := x % 10
		x /= 10
		rev = rev*10 + p
	}

	if rev >= math.MaxInt32 || rev <= math.MinInt32 {
		return 0
	}

	return rev
}
