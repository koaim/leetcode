package leetcode

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	c := make([][]rune, numRows)
	up := true
	k := 0

	for _, v := range s {
		if k == numRows-1 {
			up = false
		}

		if k == 0 {
			up = true
		}

		c[k] = append(c[k], v)

		if up {
			k++
		} else {
			k--
		}
	}

	res := make([]rune, 0)

	for _, v := range c {
		res = append(res, v...)
	}

	return string(res)
}
