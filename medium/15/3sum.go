package leetcode

import (
	"sort"
	"strconv"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(nums []int) [][]int {
	numsMap := map[int]int{}
	for i, v := range nums {
		numsMap[v] = i
	}

	result := [][]int{}
	used := map[string]bool{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if num, ok := numsMap[-sum]; ok && num != i && num != j {
				triple := []int{-sum, nums[i], nums[j]}
				key := toString(triple)

				if _, ok := used[key]; !ok {
					result = append(result, triple)
					used[key] = true
				}
			}
		}
	}

	return result
}

func toString(arr []int) string {
	sort.Ints(arr)

	return strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2])
}
