package leetcode

/*
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

Return the number that I picked.
*/

func guessNumber(n int) int {
	return guessFromTo(1, n)
}

func guessFromTo(n1, n2 int) int {
	num := n1 + (n2-n1)/2
	answer := guess(num)

	if answer == 0 {
		return num
	} else if answer == -1 {
		return guessFromTo(n1, num)
	} else {
		return guessFromTo(num+1, n2)
	}
}

func guess(int) int
