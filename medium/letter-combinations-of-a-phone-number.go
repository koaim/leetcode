package leetcode

import "unicode/utf8"

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
Return the answer in any order.
*/
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	conv := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	r := []rune(digits)
	result := conv[string(r[0])]

	for i := 1; i < utf8.RuneCountInString(digits); i++ {
		result = union(result, conv[string(r[i])])
	}

	return result
}

func union(arr1, arr2 []string) []string {
	length := len(arr1)
	for i := 0; i < length; i++ {
		for j := 0; j < len(arr2); j++ {
			arr1 = append(arr1, arr1[i]+arr2[j])
		}

	}

	return arr1[length:]
}
