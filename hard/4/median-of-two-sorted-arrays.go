package leetcode

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	l := len(nums1)
	if l == 0 {
		return float64(0)
	}

	if l%2 == 0 {
		return float64(nums1[l/2-1]+nums1[l/2]) / 2
	}

	return float64(nums1[l/2])
}
