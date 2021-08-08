package leetcode

/*Given two strings s and t, return true if t is an anagram of s, and false otherwise.*/

func isAnagram(s string, t string) bool {
	m := map[rune]int{}

	sr := []rune(s)
	tr := []rune(t)

	if len(sr) != len(tr) {
		return false
	}

	for _, v := range sr {
		m[v]++
	}

	for _, v := range tr {
		m[v]--
		if m[v] == 0 {
			delete(m, v)
		}
	}

	return len(m) == 0
}
