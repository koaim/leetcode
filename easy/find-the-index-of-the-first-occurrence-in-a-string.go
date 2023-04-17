package leetcode

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/
func strStr(haystack string, needle string) int {
	r := []rune(haystack)

	subIndex := -1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		substr := string(r[i : i+len(needle)])

		if substr == needle {
			subIndex = i
			break
		}
	}

	return subIndex
}
