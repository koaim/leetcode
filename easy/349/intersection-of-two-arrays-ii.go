package leetcode

/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.
*/
func intersection(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}

	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}

	var res []int
	for k := range set1 {
		_, ok := set2[k]
		if ok {
			res = append(res, k)
		}
	}

	return res
}
