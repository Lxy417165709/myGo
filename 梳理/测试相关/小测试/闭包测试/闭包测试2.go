package main

import (
	"fmt"
)

var (
	Ga int = 99
)



func GetGa() func() int {
	return func() int {
		Ga += 1
		return Ga
	}
}

func main() {
	b := GetGa()
	fmt.Println("main函数中：", b(), b(), b(), b())
	fmt.Println("执行闭包函数后 Ga: ", Ga)

	// main函数中： 100 101 102 103
	// 执行闭包函数后 Ga:  103
}

/*
	从全局变量Ga的结果来看，闭包只是引用了外部变量，不会创建新的变量！！！
*/