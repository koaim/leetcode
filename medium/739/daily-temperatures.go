package leetcode

/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, v := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[last] = i - last
		}
		stack = append(stack, i)
	}

	return res
}
