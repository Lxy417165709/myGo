package main

import "fmt"

func main() {

	// 1. 双引号字符串
	s1 := "abc\td"

	// 2. 反单引号字符串
	s2 := `abc\td`

	fmt.Println(s1)
	fmt.Println(s2)
	// abc	d
	// abc\td
}

/*
	总结
	1. 双引号字符串接受转义字符,单引号字符串原样输出
*/