package leetcode

/*
You are given two 0-indexed integer permutations A and B of length n.
A prefix common array of A and B is an array C such that C[i] is equal to the count of numbers that are present at or before the index i in both A and B.
Return the prefix common array of A and B.

A sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.
*/
func findThePrefixCommonArray(A []int, B []int) []int {
	result := make([]int, len(A))
	freq := map[int]int{}
	var commonCount int

	for i := 0; i < len(A); i++ {
		freq[A[i]]++
		if freq[A[i]] == 2 {
			commonCount++
		}

		freq[B[i]]++
		if freq[B[i]] == 2 {
			commonCount++
		}

		result[i] = commonCount
	}

	return result
}
