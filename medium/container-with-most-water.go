package leetcode

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
N vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
*/

import "math"

func maxArea(height []int) int {
	max := 0.0

	i, j := 0, len(height)-1

	for i <= j {
		s := math.Min(float64(height[i]), float64(height[j])) * float64(j-i)

		max = math.Max(max, s)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return int(max)
}
