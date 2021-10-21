package main

func isStrobogrammatic(num string) bool {

	length := len(num)
	l := 0
	r := length - 1
	for l <= r {
		if num[l] == 54 && num[r] == 57 || num[l] == 57 && num[r] == 54 ||
			num[l] == 56 && num[r] == 56 || num[l] == 49 && num[r] == 49 ||
			num[l] == 48 && num[r] == 48 {
			l++
			r--
		} else {
			return false
		}
	}
	return true

}
