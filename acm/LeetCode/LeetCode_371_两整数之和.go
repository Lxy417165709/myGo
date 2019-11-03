package main

import "fmt"

func getSum(a int, b int) int {
	c,d:=a,b
	for ;d!=0 ;  {
		carry:=(c&d)<<1
		c = c^d
		d = carry
		//fmt.Println(c,d)
	}
	return c

}

func main() {
	fmt.Println(getSum(-2,3))
	fmt.Println((-1)<<63)	// 不带符号位的左移
}

/*
	总结
	1. 该题是一道位运算的题目
	2. 知识点：  a + b = a^b + (a&b)<<1, 如果不能使用+号，那必须循环到(a&b)==0，此时就不用加上(a&b)<<1了
*/