package leetcode

import "fmt"

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.
*/
func addStrings(num1 string, num2 string) string {
	set := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	i, j := len(num1)-1, len(num2)-1
	var delta int
	var result []byte

	for i >= 0 || j >= 0 {
		var val1, val2 int
		if i >= 0 {
			val1 = set[num1[i]]
		}
		if j >= 0 {
			val2 = set[num2[j]]
		}

		sum := val1 + val2 + delta
		if sum > 9 {
			sum = sum - 10
			delta = 1
		} else {
			delta = 0
		}

		result = append(result, fmt.Appendf(nil, "%v", sum)...)
		i--
		j--
	}

	if delta > 0 {
		result = append(result, fmt.Appendf(nil, "%v", delta)...)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
