package main

import "fmt"

// 参考 https://blog.csdn.net/deaidai/article/details/78167367

// 返回x模2后的值
func modTwo(x int)int{
	return x & 1
}

// 消除x二进制补码中最后的1，返回消除后的值,(110)2，则返回(100)
func remove(x int) int{
	return x & (x-1)
}

// 返回x二进制补码中最后的1对应的值，(110)2，则返回(10)2
func lowbit(x int) int{
	return x & (-x)
}

func main() {
	x:=7
	fmt.Println(x,remove(x),lowbit(x))
	//fmt.Println(x^(x-1))


}

/*
	总结
	1. remove(x) = x - lowbit(x)

*/




