package main

//input: num = 12
//output: 3
// weight: 1 [1,10] 2 [2,11] 3 [3,12] [4] [5] [6] [7] [8] [9]

func GetMaxGroupNum(num int) int {
	res := 0
	max := 0
	arr := make(map[int]int)
	for i := 1; i <= num; i++ {
		value := 0
		tmp := i
		for tmp != 0 {
			value += tmp % 10
			tmp /= 10
		}
		arr[value]++
	}
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	for _, v := range arr {
		if v == max {
			res++
		}
	}
	return res
}

func main() {
	output := GetMaxGroupNum(12)
	println(output)
}
