package leetcode

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/
func generateParenthesis(n int) []string {
	result := []string{}
	generate("(", 1, 0, n, &result)
	return result
}

func generate(s string, open, closed, count int, result *[]string) {
	if open+closed == count*2 {
		*result = append(*result, s)
		return
	}

	if open == count {
		generate(s+")", open, closed+1, count, result)
	} else {
		generate(s+"(", open+1, closed, count, result)

		if open > closed {
			generate(s+")", open, closed+1, count, result)
		}
	}
}
