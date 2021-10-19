package main

import "unsafe"

const x1, y int = 1, 2
const s1 = "hello"
const (
	a1, b1      = 10, 100
	c      bool = false
)
const (
	a2 = "abc"
	b2 = len(a2)
	c2 = unsafe.Sizeof(b1)
)

//枚举
type Color int

const (
	Black Color = iota
	Red
)

func test2(color Color) {

}

const (
	Sunday = iota
	Monday
	Tuesday = 3
	Friday  = iota
)
const (
	A, B = iota, iota << 10
	C, D
)

func main() {
	const x = "xxx"
	const (
		x1 = "abc"
		x2
	)
	println(x1, x2)
	println("指针长度", c2)
	println(Friday)

	var c Color = Black
	test2(c)

	test2(1)
}
