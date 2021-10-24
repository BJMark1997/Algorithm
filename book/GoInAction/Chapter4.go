package main

import "fmt"

func test4array() {
	var array1 [5]int
	array2 := [5]int{10, 20, 30, 40, 500}
	//容量由初始化值的数量决定
	array3 := [...]int{10, 20, 30, 40, 50}
	array3[2] = 100
	fmt.Println(array1, array2, array3)
	//访问指针数组元素
	array4 := [5]*int{new(int), new(int)}
	//变量名代表整个数组
	*array4[0] = 10
	*array4[1] = 20

	var array5 [5]string
	array6 := [5]string{"Red", "Blue"}
	array5 = array6
	fmt.Println(array5)

	var array7 [3]*string
	array8 := [3]*string{new(string), new(string), new(string)}
	*array8[0] = "Red"
	*array8[1] = "Blue"
	*array8[2] = "Green"
	array7 = array8
	fmt.Printf("%v", *array7[0])

}

func main() {
	test4array()
}
