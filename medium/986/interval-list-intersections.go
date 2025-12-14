package leetcode

/*
You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. Each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.
*/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return nil
	}

	var first, second int
	var res [][]int

	for first < len(firstList) && second < len(secondList) {
		firstVal := firstList[first]
		secondVal := secondList[second]

		if secondVal[0] >= firstVal[0] && secondVal[0] <= firstVal[1] {
			res = append(res, []int{secondVal[0], min(firstVal[1], secondVal[1])})
		} else if firstVal[0] >= secondVal[0] && firstVal[0] <= secondVal[1] {
			res = append(res, []int{firstVal[0], min(firstVal[1], secondVal[1])})
		}

		if firstVal[1] < secondVal[1] {
			first++
		} else {
			second++
		}
	}

	return res
}
