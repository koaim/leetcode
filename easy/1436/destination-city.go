package leetcode

/*
You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi.
Return the destination city, that is, the city without any path outgoing to another city.

It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.
*/
func destCity(paths [][]string) string {
	type path struct {
		from bool
		to   bool
	}

	set := map[string]path{}
	for _, v := range paths {
		p := set[v[0]]
		p.from = true
		set[v[0]] = p

		pp := set[v[1]]
		pp.to = true
		set[v[1]] = pp
	}

	for k, v := range set {
		if !v.from && v.to {
			return k
		}
	}

	return ""
}
