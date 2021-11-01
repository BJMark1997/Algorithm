package main

func distributeCandies(candyType []int) int {
	set := map[int]struct{}{}
	for _, t := range candyType {
		set[t] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}
