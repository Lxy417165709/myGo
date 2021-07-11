package main

import "fmt"

// 全局变量
var a, b, c = 1, false, "hello"

func funciton(num int){
	// num是形式参数
	fmt.Println(num)
}

func main(){
	// 局部变量
	var d, e, f = 1, false, "hello"
	fmt.Println(a,b,c,d,e,f)
	// 块作用域
	{
		var e = 1
		fmt.Println("里面的e",e)
	}
	fmt.Println("外面的e",e)
}

/*
	总结
	1. go作用域规则和C++一样
	2. 子能访问父的变量,父一般不能访问子的变量(函数内可以访问全局变量,函数外无法访问函数内局部变量无法访问(一般))
 */
