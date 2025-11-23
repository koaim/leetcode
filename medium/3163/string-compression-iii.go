package leetcode

import (
	"fmt"
	"strings"
)

/*
Given a string word, compress it using the following algorithm:

		-Begin with an empty string comp. While word is not empty, use the following operation:
	    	- Remove a maximum length prefix of word made of a single character c repeating at most 9 times.
	    	- Append the length of the prefix followed by c to comp.

Return the string comp.
*/
func compressedString(word string) string {
	var sb strings.Builder
	prev := word[0]
	buf := 1

	for i := 1; i < len(word); i++ {
		if word[i] == prev && buf < 9 {
			buf++
		} else {
			sb.WriteString(fmt.Sprintf("%v", buf))
			sb.WriteString(string(prev))
			buf = 1
			if word[i] != prev {
				prev = word[i]
			}
		}
	}

	sb.WriteString(fmt.Sprintf("%v", buf))
	sb.WriteString(string(prev))

	return sb.String()
}
