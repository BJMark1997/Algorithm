package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

var v1 int
var v2 int
var v3 [10]int
var v4 []int
var v5 struct {
	f int
}
var v6 *int
var v7 map[string]int
var v8 func(a int) int

var (
	v9  int
	v10 string
)

func myInit() {
	var v1 int = 10
	var v2 = 20
	v3 := 30
	v1, v2 = v2, v1
	fmt.Println(v1, v2, v3)
}

func GetName() (firstname, lastname, nickname string) {
	return "May", "Chan", "Chibi Maruko"
}

const Pi float64 = 3.14
const zero = 0.0
const (
	size int64 = 1024
	eof        = -1
)
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"
const mask = 1 << 3

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

func myBool() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Println(v1, v2)
}

//IsEqual 浮点数比较
func IsEqual(f1, f2, p float64) bool {
	p = 0.01
	return math.Abs(f1-f2) < p
}

//复数
var val1 complex64

//myString字符串
func myString() {
	str := "Hello World"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c", str, ch)
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
	//Unicode
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func myRune() {
	str := "Hello World"
	decodeRune, s := utf8.DecodeRune([]byte(str))
	fmt.Println(string(decodeRune), s)
}

//myArray 数组
func myArray() {
	st1 := [2]struct{ x, y int32 }{}
	st1[1].x = 10
	st1[1].y = 20
	st2 := [1000]*float64{}
	st3 := [3][5]int{}
	st4 := [2][2][2]float64{}
	fmt.Println(st1, st2, st3, st4)
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println(array)
}

func mySlice() {
	myslice1 := make([]int, 5)
	myslice2 := make([]int, 5, 10)
	myslice3 := []int{1, 2, 3}
	ints := append(myslice3, 1)
	output := append(ints, myslice2...)
	fmt.Println(myslice1, myslice2, myslice3, output)
}
func main() {
	myInit()
	_, _, nickname := GetName()
	fmt.Println(nickname)

	val1 = 3.2 + 12i
	val3 := complex(3.2, 12)
	fmt.Println(val1, val3)

	//myString()
	myRune()
	myArray()

	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println(array)

	myarray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := myarray[:5]
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	for _, v := range myarray {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
