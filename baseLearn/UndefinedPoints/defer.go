package main

import "fmt"
func myFunction0() int {
	i := 0
	defer func() {
		i = i + 2
	}()
	return i
}
func myFunction1() (z int){
	i := 0
	defer func() {
		z = i + 2
	}()
	return i
}
func myFunction2() (z int){
	i := 0
	defer func() {
		z = i + 2
	}()
	return 100
}
func myFunction3() (z int){
	i := 0
	defer func() {
		z = i + 2
	}()
	return
}

func main() {
	fmt.Println(myFunction0())
	fmt.Println(myFunction1())
	fmt.Println(myFunction2())
	fmt.Println(myFunction3())
	// 0
	// 2
	// 2
	// 2

	// myFunction2()和myFunction3()的返回值由defer里面的函数决定(因为defer中操作了返回值 z)



	fmt.Println("打开文件1")
	defer func() {
		fmt.Println("关闭文件1")
	}()
	fmt.Println("打开文件2")
	defer func() {
		fmt.Println("关闭文件2")
	}()
	fmt.Println("打开文件3")
	defer func() {
		fmt.Println("关闭文件3")
	}()

	// 打开文件1
	// 打开文件2
	// 打开文件3
	// 关闭文件3
	// 关闭文件2
	// 关闭文件1
}
/*
	总结
	1. defer和return整体执行顺序
			- 先return
			- 再defer (defer也可以决定返回值)
			- 最后跳出函数
	2. 一个函数内的defer执行顺序是栈的顺序(先定义的defer后执行)

*/