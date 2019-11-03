package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. 初始化
	arr := [5]int {1,2,3,4,5}
	sli := arr[0:3]
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc042068060 0xc042068060
	// [1 2 3 4 5] [1 2 3]
	// 3 5 24
	// --- 切片容量等于数组的长度,切片长度等于取出数组元素的个数 ---

	// 2. 把切片中index==1的元素改为8
	sli[1]=8
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc042068060 0xc042068060
	// [1 8 3 4 5] [1 8 3]
	// 3 5 24
	// --- 可以发现切片变了,数组也变了(因为现在切片中的指针变量指向这个数组的首地址)~ ---

	// 3. 向切片追加元素
	sli = append(sli,10)
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc04200a300 0xc04200a300
	// [1 8 3 7 5] [1 8 3 7]
	// 4 5 24
	// --- 切片指针不变,数组元素变了 ---

	// 4. 向切片继续追加元素
	sli = append(sli,11)
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc042068060 0xc042068060
	// [1 8 3 10 11] [1 8 3 10 11]
	// 5 5 24
	// --- 切片指针不变,数组元素变了 ---

	// 5. 向切片继续追加元素
	sli = append(sli,12)
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc042068060 0xc042078050
	// [1 8 3 10 11] [1 8 3 10 11 12]
	// 6 10 24
	// --- 可以发现切片容量和切片指针变了~ ---

	// 6. 把切片中index==1的元素改为1000
	sli[1]=100
	fmt.Printf("%p %p\n",&arr,sli)
	fmt.Println(arr,sli)
	fmt.Println(len(sli),cap(sli),unsafe.Sizeof(sli))
	// 0xc042068060 0xc042078050
	// [1 8 3 10 11] [1 100 3 10 11 12]
	// 6 10 24
	// --- 可以发现切片变了,数组没变(因为现在切片中的指针变量不指向这个数组的首地址)~ ---

}
/*
	总结
	1. 切片中有一个指针,指向数组首元素的地址~
	2. 切片其实是一个结构体,里面有一个64位指针变量,int64长度变量,int64容量变量
	3. 当用数组初始化切片时,切片容量等于数组的长度,切片长度等于取出数组元素的个数

*/