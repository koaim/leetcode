package leetcode

import "strings"

/*
You are given an absolute path for a Unix-style file system, which always begins with a slash '/'.
Your task is to transform this absolute path into its simplified canonical path.
*/
func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	s := &Stack{}

	for _, v := range pathArr {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			s.Pop()
		} else {
			s.Push(v)
		}
	}

	return "/" + strings.Join(s.val, "/")
}

type Stack struct {
	val []string
}

func (s *Stack) Push(v string) {
	s.val = append(s.val, v)
}

func (s *Stack) Pop() {
	if len(s.val) == 0 {
		return
	}

	s.val = s.val[:len(s.val)-1]
}

func (s *Stack) Len() int {
	return len(s.val)
}
