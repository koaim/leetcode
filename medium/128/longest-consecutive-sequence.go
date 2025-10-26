package leetcode

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := map[int]struct{}{}
	for _, v := range nums {
		numsMap[v] = struct{}{}
	}

	cons := map[int]int{}
	for _, v := range nums {
		_, ok := cons[v]
		if ok {
			continue
		}

		_, ok = numsMap[v-1]
		if ok {
			continue
		}

		next := v + 1
		for {
			_, ok := numsMap[next]
			if ok {
				next++
			} else {
				cons[v] = next - 1
				break
			}
		}
	}

	var max int
	for k, v := range cons {
		length := v - k
		if length > max {
			max = length
		}
	}

	return max + 1
}
