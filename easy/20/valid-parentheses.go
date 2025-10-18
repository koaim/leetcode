package leetcode

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
*/

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	st := Stack{}
	valid := true

	for _, v := range s {
		if v == ')' || v == ']' || v == '}' {
			parent := st.Get()
			if parent == nil || *parent != m[v] {
				valid = false
			}
		} else {
			st.Add(v)
		}
	}

	return valid && st.Len() == 0
}

type Stack struct {
	val []rune
}

func (s *Stack) Add(v rune) {
	s.val = append(s.val, v)
}

func (s *Stack) Get() *rune {
	if len(s.val) == 0 {
		return nil
	}

	v := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]

	return &v
}

func (s *Stack) Len() int {
	return len(s.val)
}
