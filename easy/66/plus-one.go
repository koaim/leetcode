package leetcode

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.
*/
func plusOne(digits []int) []int {
	i := len(digits) - 1
	for {
		if digits[i] == 9 && i != 0 {
			digits[i] = 0
			i--
		} else if digits[i] < 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
			digits = append([]int{1}, digits...)
			break
		}
	}

	return digits
}
