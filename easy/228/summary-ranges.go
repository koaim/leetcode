package leetcode

import "fmt"

/*
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
*/
func summaryRanges(nums []int) []string {
	var res []string
	var start, end *int
	var prev int

	for _, v := range nums {
		if start == nil {
			start = &v
		} else if (v - prev) == 1 {
			end = &v
		} else {
			if end != nil {
				res = append(res, fmt.Sprintf("%v->%v", *start, *end))
			} else {
				res = append(res, fmt.Sprintf("%v", *start))
			}
			start, end = &v, nil
		}
		prev = v
	}

	if start != nil && end != nil {
		res = append(res, fmt.Sprintf("%v->%v", *start, *end))
	} else if start != nil {
		res = append(res, fmt.Sprintf("%v", *start))
	}

	return res
}
