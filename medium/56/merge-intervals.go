package leetcode

/*
Given an array of intervals where intervals[i] = [starti, endi].
Merge all overlapping intervals and return an array of the non-overlapping intervals that cover all the intervals in the input.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
*/

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		ps := intervals[i-1][0]
		pe := intervals[i-1][1]

		cs := intervals[i][0]
		ce := intervals[i][1]

		var union [][]int
		if ps <= cs && cs <= pe {
			union = [][]int{{ps, max(ce, pe)}}
		} else if cs <= ps && ps <= ce {
			union = [][]int{{cs, max(ce, pe)}}
		} else if cs <= ps && ce <= pe {
			union = [][]int{{ps, pe}}
		}

		if len(union) > 0 {
			intervals = append(intervals[:i-1], intervals[i+1:]...)
			intervals = append(intervals[:i-1], append(union, intervals[i-1:]...)...)
			i--
		}
	}

	return intervals
}
