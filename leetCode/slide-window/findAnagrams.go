package slide_window

func findAnagrams(s string, p string) []int {
	var ans []int
	window := make(map[byte]int, 0)
	tmp := make(map[byte]int, 0)

	ls, lp := len(s), len(p)
	if lp > ls {
		return ans
	}

	for _, v := range p {
		if _, ok := tmp[byte(v)]; !ok {
			tmp[byte(v)] = 1
		} else {
			tmp[byte(v)]++
		}
	}

	left, right, valid := 0, 0, 0
	for right < ls {
		r := s[right]
		right++
		if _, ok := tmp[r]; ok {
			window[r]++
			if window[r] == tmp[r] {
				valid++
			}
		}
		for right-left >= lp {
			if valid == len(tmp) {
				ans = append(ans, left)
			}
			lc := s[left]
			left++
			if _, ok := tmp[lc]; ok {
				if window[lc] == tmp[lc] {
					valid--
				}
				window[lc]--
			}
		}

	}
	return ans
}
