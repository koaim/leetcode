package leetcode

/*
Given n points on a 2D plane, find if there is such a line parallel to the y-axis that reflects the given points symmetrically.
In other words, answer whether or not if there exists a line that after reflecting all points over the given line, the original points' set is the same as the reflected ones.

Note that there can be repeated points.
*/
func isReflected(points [][]int) bool {
	minX, maxX := points[0][0], points[0][0]
	set := map[[2]int]struct{}{}

	for _, v := range points {
		x := v[0]
		y := v[1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		set[[2]int{x, y}] = struct{}{}
	}

	axis := float64(minX+maxX) / 2.0

	for _, v := range points {
		reflected := reflectedPoint(v, axis)
		if _, ok := set[reflected]; !ok {
			return false
		}
	}

	return true
}

func reflectedPoint(point []int, xSim float64) [2]int {
	x := int((2 * xSim) - float64(point[0]))
	y := point[1]

	return [2]int{x, y}
}
