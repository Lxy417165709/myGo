package main

import "fmt"

func demo (name string,hobby ...string){
	fmt.Println("我是",name)
	fmt.Println("喜欢")
	for i,s := range hobby{
		fmt.Println(i,s)
	}
	fmt.Printf("%T\n",hobby)
}

func main() {
	demo("蔡徐坤","唱","跳","rap","篮球")
	// 我是 蔡徐坤
	// 喜欢
	// 0 唱
	// 1 跳
	// 2 rap
	// 3 篮球
	// []string
}


/*
	总结
	1. 可变参数的本质是一个切片
	2. 可变参数只能写在函数参数列表的最后
*/
