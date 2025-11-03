package leetcode

/*
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself.
A divisor of an integer x is an integer that can divide x evenly.

Given an integer n, return true if n is a perfect number, otherwise return false.
*/
func checkPerfectNumber(num int) bool {
	var divisors []int

	for i := 1; i < (num/2)+1; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}

	var sum int
	for _, v := range divisors {
		sum += v
	}

	return sum == num
}
