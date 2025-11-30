package leetcode

func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}

	var isoString []byte
	replacement := map[byte]byte{}
	set := map[byte]struct{}{}

	for i := 0; i < len(s); i++ {
		v, ok := replacement[s[i]]
		if ok {
			isoString = append(isoString, v)
		} else {
			if _, ok := set[t[i]]; ok {
				return false
			}

			replacement[s[i]] = t[i]
			set[t[i]] = struct{}{}
			isoString = append(isoString, t[i])
		}
	}

	return string(isoString) == t
}
