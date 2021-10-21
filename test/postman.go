package test

func postman(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ret := 0
	index := 0
	for {
		if index >= len(arr) {
			break
		}
		if arr[index] == 0 {
			index++
		} else {
			ret++
			if index+3 < len(arr) {
				index = index + 3
			} else {
				break
			}
		}
	}
	return ret
}
