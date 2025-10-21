package leetcode

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.
*/
func mySqrt(x int) int {
	var res int

	i := 1
	for {
		if i*i < x {
			res = i
			i++
		} else if i*i == x {
			res = i
			break
		} else {
			break
		}
	}

	return res
}
