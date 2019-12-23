package main

import "fmt"

func main() {
	// 1. 带参数及返回值的匿名函数
	ans := func(a int,b int) int {
		return a + b
	}(10,20)
	fmt.Println(ans)
	// 30

	// 2. 函数变量
	f := func(a int,b int) int {
		return a + b
	}
	fmt.Println(f(5,20))
	fmt.Printf("%T\n",f)
	// 25
	// func(int, int) int


}

/*
	总结
	1. 函数变量不用(),而匿名函数调用要()	看上面就可以了...好难说
	2. 函数也是一种变量类型~
*/