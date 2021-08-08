package leetcode

/*
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, v := range strs {
		sv := sortStr(v)
		m[sv] = append(m[sv], v)
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func sortStr(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
