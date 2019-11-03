package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcde李同学"

	// 1. 字节长度及字符长度 (对于UTF-8编码,汉字字符占3个字节,英文字符占1个字节)
	byteLen := len(s)
	s1:= []rune(s)
	charLen := len(s1)
	fmt.Printf("%T %T\n",s,s1)
	fmt.Println("字节长度: ",byteLen)
	fmt.Println("字符长度: ",charLen)
	// string []int32
	// 字节长度:  14
	// 字符长度:  8

	// 2. 截取字符串
	ltx := s[5:]
	fmt.Println(ltx)
	fmt.Printf("%T\n",ltx)
	// 李同学
	// string

	// 3. 常用函数
	x := "123abbaawwc李wecdde"

	// 3.10 第一次出现的索引
	fmt.Println(strings.Index(x,"a"))
	// 3.11 最后一次出现的索引
	fmt.Println(strings.LastIndex(x,"a"))
	// 3
	// 7

	// 3.20 是否以指定字符串开头
	fmt.Println(strings.HasPrefix(x,"123"))
	// 3.21 是否以指定字符串结尾
	fmt.Println(strings.HasSuffix(x,"李wecdde"))
	// 3.22 是否包含指定字符串
	fmt.Println(strings.Contains(x,"李"))
	// true
	// true
	// true

	// 3.30 转为小写
	fmt.Println(strings.ToLower(x))
	// 3.31 转为大写
	fmt.Println(strings.ToUpper(x))
	// 123abbaawwc李wecdde
	// 123ABBAAWWC李WECDDE

	// 3.40 字符串替换(n为替换次数,n<0表示全部替换)
	fmt.Println(strings.Replace(x,"李","王",-2))
	// 3.41 字符串重复
	fmt.Println(strings.Repeat(x,2))
	// 123abbaawwc王wecdde
	// 123abbaawwc李wecdde123abbaawwc李wecdde

	// 3.50 去除字符串前后所有指定字符
	fmt.Println(strings.Trim(x,"1"))
	fmt.Println(strings.Trim(x," "))
	fmt.Println(strings.TrimSpace(x))	// 是上面的简写
	// 23abbaawwc李wecdde
	// 123abbaawwc李wecdde
	// 123abbaawwc李wecdde

	// 3.60 字符串切割
	fmt.Println(strings.Split(x,"a"))
	// 3.61 字符串合并
	sli := []string{"123","456","789"}
	fmt.Println(strings.Join(sli,"-"))
	// [123 bb  wwc李wecdde]
	// 123-456-789



}


/*
	总结
	1. 看上面就懂的啦,你这么聪明~~~
*/