package leetcode

/*
You are given an array representing a row of seats where seats[i] = 1 represents a person sitting in the ith seat, and seats[i] = 0 represents that the ith seat is empty (0-indexed).
There is at least one empty seat, and at least one person sitting.
Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to the closest person.
*/
func maxDistToClosest(seats []int) int {
	prevOne := -1
	dist := 0

	for i, v := range seats {
		if v == 0 && i != len(seats)-1 {
			continue
		} else if v == 0 && i == len(seats)-1 {
			dist = max(dist, i-prevOne)
		} else if v == 1 && i == 0 {
			prevOne = i
			continue
		} else {
			if prevOne == -1 {
				dist = max(dist, i)
			} else {
				dist = max((i-prevOne)/2, dist)
			}
			prevOne = i
		}
	}

	return dist
}
