package main

import "fmt"

func main() {
	fmt.Println("开始了开始了~")
	goto newPoint
	fmt.Println("我是不会执行的!")

	newPoint: {
		fmt.Println("我执行了!")
	}
	// 开始了开始了~
	// 我执行了!
}

/*
	总结
	1. goto可以让代码跳到指定位置执行
	2. goto不利于代码的结构 (所以不推荐用)
*/