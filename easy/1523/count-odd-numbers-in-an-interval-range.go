package leetcode

/*
Given two non-negative integers low and high. 
Return the count of odd numbers between low and high (inclusive).
*/
func countOdds(low int, high int) int {
	var result int

	for i := low; i <= high; i++ {
		if i%2 != 0 {
			result++
		}
	}

	return result
}
