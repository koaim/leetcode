package leetcode

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
    - answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
    - answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

Note that the integers in the lists may be returned in any order.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := map[int]struct{}{}
	map2 := map[int]struct{}{}
	result := make([][]int, 2)

	for _, v := range nums1 {
		map1[v] = struct{}{}
	}

	for _, v := range nums2 {
		map2[v] = struct{}{}
	}

	for k := range map1 {
		_, ok := map2[k]
		if !ok {
			result[0] = append(result[0], k)
		}
	}

	for k := range map2 {
		_, ok := map1[k]
		if !ok {
			result[1] = append(result[1], k)
		}
	}

	return result
}
