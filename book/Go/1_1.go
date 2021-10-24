package main

var x int
var f float32 = 1.6
var s = "abc"

func test() (int, string) {
	return 1, "abc"
}
func test1() {
	s := "abc"
	println(&s)
	s, y := "hello", 30
	println(&s, y)
	{
		s, z := "hello", 20
		println(&s, z)
	}
}

func main() {
	x := 123
	_ = x
	var y, z int
	_, _ = y, z
	var (
		a int
		_ = a
		b int
		_ = b
	)

	n, s := 0x1234, "Hello,World!"
	println(x, s, n)
	//多变量赋值,先计算,再从左到右依次赋值
	data, i := []int{0, 1, 2}, 0
	i, data[i] = 2, 100

	_, s = test()
	println(s)

	test1()
}
