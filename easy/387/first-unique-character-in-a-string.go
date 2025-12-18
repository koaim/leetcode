package leetcode

/*
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
*/
func firstUniqChar(s string) int {
	set := map[rune]int{}
	for _, v := range s {
		set[v]++
	}

	for i, v := range s {
		if set[v] == 1 {
			return i
		}
	}

	return -1
}
