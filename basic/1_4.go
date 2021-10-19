package main

type Point struct {
}
type Point1 struct {
}

func main() {
	a := []int{0, 0, 0}
	a[1] = 10

	b := make([]int, 3)
	b[1] = 10

	c := new([]int)
	(*c)[1] = 10

	var d byte = 100
	var n int = int(d)
	_ = n

	p := &Point{}
	p2 := (*Point1)(p)
	_ = p2

	ints := make(chan int)
	i := (<-chan int)(ints)
	_ = i
}
