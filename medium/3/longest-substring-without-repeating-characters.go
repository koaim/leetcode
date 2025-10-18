package leetcode

/*
Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {
	c := make(map[rune]int, 0)

	max := 0

	for j, v := range s {
		i, ok := c[v]
		if !ok {
			c[v] = j
		} else {
			if len(c) > max {
				max = len(c)
			}

			if i == j-1 {
				c = make(map[rune]int, 1)
				c[v] = j
			} else {
				c[v] = j
				for a, b := range c {
					if b < i {
						delete(c, a)
					}
				}
			}
		}
	}

	if len(c) > max {
		max = len(c)
	}

	return max
}
