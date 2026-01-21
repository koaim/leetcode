package leetcode

/*
Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array.
The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:
    - |a - x| < |b - x|, or
    - |a - x| == |b - x| and a < b
*/
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	var res []int
	for i := left; i < left+k; i++ {
		res = append(res, arr[i])
	}

	return res
}
