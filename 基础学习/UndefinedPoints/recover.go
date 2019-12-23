package main

import "fmt"

func fun1(){
	fmt.Println("我是fun1前半部分")
	fun2()
	fmt.Println("我是fun1后半部分")
}
func fun2(){
	defer func() {
		recover()
	}()
	fmt.Println("我是fun2前半部分")
	fun3()
	defer func() {
		fmt.Println("我是不会执行的了~呜呜呜呜")
		recover()
	}()
	fmt.Println("我是fun2后半部分")
}
func fun3(){
	fmt.Println("我是fun3前半部分")
	panic("我fun3出错了~~~")
	fmt.Println("我是fun3后半部分")
}

func main() {
	fmt.Println("main前半部分")
	fun1()
	fmt.Println("main后半部分")

	// (无recover时)
	// main前半部分
	// 我是fun1前半部分
	// 我是fun2前半部分
	// 我是fun3前半部分
	// panic: 我fun3出错了~~~

	// (在fun2中加入recover后)
	// main前半部分
	// 我是fun1前半部分
	// 我是fun2前半部分
	// 我是fun3前半部分
	// 我是fun1后半部分
	// main后半部分

}
/*
	总结
	1. panic在函数之间有向上传递性,直到这些函数里面有recover,panic才算被消除
       (处理后就表示程序正常了,下面的代码可以执行了)
	2. 若要恢复异常,recover应该定义在defer里面,且defer代码要写在panic产生的代码之前
 */