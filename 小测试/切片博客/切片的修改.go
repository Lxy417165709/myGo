package main

import (
	"fmt"

)

func main() {
	//array := [5]int{1,2,3,4,5}
	//slice := array[:]
	//
	//fmt.Println("修改切片之前:")
	//fmt.Println("	切片:",slice)
	//fmt.Println("	数组:",array)
	//
	//slice[0] = 8	// 把切片元素0的值改为8
	//fmt.Println("修改切片之后:")
	//fmt.Println("	切片:",slice)
	//fmt.Println("	数组:",array)
	//
	//// 输出
	//// 修改切片之前:
	////	切片: [1 2 3 4 5]
	////	数组: [1 2 3 4 5]
	////修改切片之后:
	////	切片: [8 2 3 4 5]
	////	数组: [8 2 3 4 5]


	array := [5]int{1,2,3,4,5}
	slice := array[:]

	fmt.Println("修改数组之前:")
	fmt.Println("	切片:",slice)
	fmt.Println("	数组:",array)

	array[0] = -8	// 把数组元素0的值改为8
	fmt.Println("修改数组之后:")
	fmt.Println("	切片:",slice)
	fmt.Println("	数组:",array)

	// 输出
	// 修改数组之前:
	//	切片: [1 2 3 4 5]
	//	数组: [1 2 3 4 5]
	// 修改数组之后:
	//	切片: [-8 2 3 4 5]
	//	数组: [-8 2 3 4 5]
}
