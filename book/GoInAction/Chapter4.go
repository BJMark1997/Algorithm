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

	var array9 [4][2]int
	array9 = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	array10 := [4][2]int{1: {0, 1}}
	var array11 = array9[1]
	fmt.Println(array9, array10, array11)

	slice := make([]string, 5)
	slice1 := make([]string, 3, 5)
	slice2 := []string{"Redis", "Docker"}
	var slice3 []int
	slice4 := []int{10, 20, 30, 40, 50}
	slice4[1] = 25
	slice5 := slice4[1:3]
	fmt.Println(slice, slice1, slice2, slice3, slice5)
	source := slice4[1:2:3]
	println(len(source), cap(source))
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	new := append(s1, s2...)
	println(new)

	for index, value := range new {
		fmt.Printf("Index: %d Value %d\n", index, value)
	}

	newSlice := [][]int{{10}, {100, 200}, {1, 2, 3}}
	newSlice[0] = append(newSlice[0], 10)

	colors := map[string]string{}
	colors["Red"] = "#dal1337"
	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}

	colors1 := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral":     "#ff7F50",
	}
	for key, value := range colors1 {
		fmt.Printf("Key %s Value: %s\n", key, value)
	}
	delete(colors, "Coral")

}

func foo(slice []int) []int {
	slice = make([]int, 10)
	return slice
}
func main() {
	test4array()
}
