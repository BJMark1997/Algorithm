package main

import "fmt"

//字符串是不可变类型,内部用指针指向UTF-8字节数组
//不能用&获取字节元素指针
//不可变类型,无法修改字节数组 尾部不含NULL

func main() {
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
	//不能转义
	s1 := `a
    b\r\n\x00
    c`
	s2 := "hello,world"
	s3 := "h" + "2"
	s4 := s2[1:4]
	println(s1)
	println(s2, s3)
	println(s4)
	//单引号字符常量表示Unicode Code Point 支持 \uFFFF \u7FFFFFFFF \xFF 对应rune
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c1), c2)

	//修改字符串
	s7 := "abcd"
	bs := []byte(s7)
	bs[1] = 'B'
	s5 := string(bs)
	println(s5)

	s6 := "电脑"
	runes := []rune(s6)
	println(string(runes))

	//遍历
	s8 := "abc汉字"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s8[i])
	}
	fmt.Println()

	for _, r := range s8 {
		fmt.Printf("%c,", r)
	}
}
