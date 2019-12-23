package main

import "fmt"

// 参数变量i和内部函数形成一个闭包
func father(i int) func() int{
	return func() int {
		i++
		return i
	}

}


func main() {
	f1:=father(0)
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	// 1
	// 2
	// 3

	f2:=father(-50)
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	// -49
	// -48
	// -47

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	// 4
	// 5
	// 6
}

/*
	总结(闭包概述)
	1. 闭包不是Go语言独有的概念,许多编程编程语言中都有闭包
	2. 闭包就是解决局部变量不能被外部访问的一种解决方案
	3. 闭包是把函数作为返回值的一种应用
	4. 闭包不单是一个函数,应该是函数+形参/局部变量
*/