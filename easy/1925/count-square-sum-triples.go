package leetcode

import "math"

/*
A square triple (a,b,c) is a triple where a, b, and c are integers and a2 + b2 = c2.

Given an integer n, return the number of square triples such that 1 <= a, b, c <= n.
*/
func countTriples(n int) int {
	var result int

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			square := i*i + j*j
			if square <= n*n && math.Mod(math.Sqrt(float64(square)), 1.0) == 0 {
				result++
			}
		}
	}

	return result
}
