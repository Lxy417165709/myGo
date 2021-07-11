package main

import "fmt"

/*
	一个slice是一个数组某个部分的引用。
	在内存中，它是一个包含3个域的结构体：指向slice中第一个元素的指针，slice的长度，以及slice的容量。
	长度是下标操作的上界，如x[i]中i必须小于长度。容量是分割操作的上界，如x[i:j]中j不能大于容量。

	来源: https://www.w3cschool.cn/go_internals/go_internals-cyz6282h.html
*/
func main() {
	slice := make([]int,3,10)
	slice[0],slice[1],slice[2] = 0,1,2
	fmt.Println("slice[2]: ", slice[2])
	fmt.Println("slice[1:5]: ",slice[1:5])

	// 输出
	// slice[2]:  2
	// slice[1:5]:  [1 2 0 0]

	// 所以，上面的说法是正确的。
}

