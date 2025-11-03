package leetcode

import "slices"

/*
Given an array of intervals where intervals[i] = [starti, endi].
Merge all overlapping intervals and return an array of the non-overlapping intervals that cover all the intervals in the input.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
*/
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(v1, v2 []int) int {
		if v1[0] > v2[0] {
			return 1
		}
		if v1[0] == v2[0] && v1[1] >= v2[1] {
			return 1
		}

		return -1
	})

	var i int
	for i != len(intervals)-1 {
		if intervals[i+1][0] >= intervals[i][0] && intervals[i+1][0] <= intervals[i][1] {
			intervals[i] = []int{intervals[i][0], max(intervals[i][1], intervals[i+1][1])}
			intervals = slices.Delete(intervals, i+1, i+2)
		} else {
			i++
		}
	}

	return intervals
}
