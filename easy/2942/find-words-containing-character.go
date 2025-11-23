package leetcode

/*
You are given a 0-indexed array of strings words and a character x.
Return an array of indices representing the words that contain the character x.

Note that the returned array may be in any order.
*/
func findWordsContaining(words []string, x byte) []int {
	var result []int
	for i, v := range words {
		if contains(v, x) {
			result = append(result, i)
		}
	}

	return result
}

func contains(s string, x byte) bool {
	for _, v := range s {
		if string(v) == string(x) {
			return true
		}
	}

	return false
}
