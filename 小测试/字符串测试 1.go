package main

import "fmt"

func main() {
	a := "abc"
	fmt.Printf("%p %s\n",&a,a)
	b := "abc"
	fmt.Printf("%p %s\n",&b,b)

	// 0xc0420481c0 abc
	// 0xc0420481e0 abc

	a = a + b
	fmt.Printf("%p %s\n",&a,a)
	// 0xc04203a1d0 abcabc

	for i:=0;i<10 ;i++  {
		a = a + a
	}
	fmt.Printf("%p %s\n",&a,a)

	c := a
	fmt.Printf("%p %s\n",&c,c)
}

/*
	测试结果
	1. 字符串传递是值传递 (会开辟新的内存空间)
	2. 字符串的追加(使用"+"号),不会改变原字符串的内存地址
*/