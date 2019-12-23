package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1.字符串转十进制整数
	s1 := "10"
	// 1.0 strconv.ParseInt
	i1, _ := strconv.ParseInt(s1, 10, 64)
	// 1.1 strconv.Atoi (简写)
	i2, _ := strconv.Atoi(s1)
	fmt.Println(i1, i2)
	// 10 10

	// 2. 字符串转十进制浮点数
	s2 := "10.5"
	f1, _ := strconv.ParseFloat(s2, 64)
	fmt.Println(f1)
	// 10.5

	// 3. 十进制整数转字符串
	i := 19
	// 3.1 strconv.FormatInt
	s3 := strconv.FormatInt(int64(i), 10)
	// 3.2 strconv.Itoa
	s4 := strconv.Itoa(i)
	fmt.Println(s3, s4)
	// 19 19

	// 4. 十进制浮点数转字符串
	f2 := 1059.78444
	s5 := strconv.FormatFloat(f2, 'f', 2, 64)
	fmt.Println(s5)
	// 1059.78

}

/*
	总结
	1. 	strconv.ParseXXX		为把其他类型转为字符串 (对于十进制整数有strconv.Atoi方法简写)
		strconv.FormatXXX		为把字符串转为其他类型 (对于十进制整数有strconv.Itoa方法简写)
	2.  详细情况请查看官方文档	https://studygolang.com/pkgdoc
*/
