package main

import "fmt"

func main() {
	recover()
	fmt.Println("我在正常执行~")
	defer func() {
		fmt.Println("我是一个defer~~~")
	}()
	panic("出错了~呜呜呜")
	fmt.Println("看看我还能不能执行(答案是不能执行了)")
	defer func() {
		fmt.Println("这个defer不能执行")
	}()

	// 我在正常执行~
	// 我是一个defer~~~
	// panic: 出错了~呜呜呜
}

/*
	总结
	1. panic后面的代码不能执行(在没有recover的情况下),即使是defer代码
	2. defer在panic出现后也会执行(只要defer代码写在panic前),defer先于panic执行
	3. 代码执行顺序:return defer panic

 */