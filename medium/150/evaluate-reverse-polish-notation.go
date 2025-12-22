package leetcode

import "strconv"

/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
Evaluate the expression. Return an integer that represents the value of the expression.
*/
func evalRPN(tokens []string) int {
	s := Stack{}

	for _, v := range tokens {
		if isOperator(v) {
			a := s.Pop()
			b := s.Pop()

			switch v {
			case "+":
				s.Push(a + b)
			case "-":
				s.Push(b - a)
			case "*":
				s.Push(a * b)
			default:
				s.Push(b / a)
			}
		} else {
			vInt, _ := strconv.Atoi(v)
			s.Push(vInt)
		}
	}

	return s.Pop()
}

type Stack struct {
	val []int
}

func (s *Stack) Push(v int) {
	s.val = append(s.val, v)
}

func (s *Stack) Pop() int {
	val := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]

	return val
}

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}

	return false
}