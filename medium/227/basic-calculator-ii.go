package leetcode

import "unicode"

/*
Given a string s which represents an expression, evaluate this expression and return its value.
The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].
*/
func calculate(s string) int {
	stack := &Stack{}
	prev := "+"
	num := 0

	for i, v := range s {
		if unicode.IsDigit(v) {
			num = num*10 + int(v-'0')
		} else if v == ' ' {
			continue
		} else if i == len(s)-1 || v == '+' || v == '-' || v == '*' || v == '/' {
			if prev == "+" {
				stack.Push(num)
			} else if prev == "-" {
				stack.Push(-num)
			} else if prev == "*" {
				val := stack.Pop()
				stack.Push(val * num)
			} else if prev == "/" {
				val := stack.Pop()
				stack.Push(val / num)
			}

			prev = string(v)
			num = 0
		}
	}

	sum := 0
	for stack.Len() != 0 {
		sum += stack.Pop()
	}

	return sum
}

type Stack struct {
	val []int
}

func (s *Stack) Push(v int) {
	s.val = append(s.val, v)
}

func (s *Stack) Pop() int {
	if len(s.val) == 0 {
		return 0
	}
	res := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return res
}

func (s *Stack) Len() int {
	return len(s.val)
}
