package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type data struct {
		a int
	}
	d := data{a: 1234}
	var p1 *data
	p1 = &d
	fmt.Printf("%p,%v\n", p1, p1.a)
	//不能对指针进行加减法运算

	x := 0x12345678
	p2 := unsafe.Pointer(&x)
	n := (*[4]byte)(p2)

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}

}
