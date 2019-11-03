package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//slice1 := []int{1,2,3}
	//fmt.Println("slice1的内容: ",slice1)
	//fmt.Println("			内存地址: ",unsafe.Pointer(&slice1))
	//fmt.Printf("			指向的内存地址: %p\n",slice1)
	//fmt.Printf("			切片的长度: %d\n", len(slice1))
	//fmt.Printf("			切片的容量: %d\n",cap(slice1))
	//
	//slice2:= slice1
	//fmt.Println("slice2的内容: ",slice2)
	//fmt.Println("			内存地址: ",unsafe.Pointer(&slice2))
	//fmt.Printf("			指向的内存地址: %p\n",slice2)
	//fmt.Printf("			切片的长度: %d\n", len(slice2))
	//fmt.Printf("			切片的容量: %d\n",cap(slice2))
	//
	//// 输出
	//// slice1的内容:  [1 2 3]
	////			内存地址:  0xc042002440
	////			指向的内存地址: 0xc04200c360
	////			切片的长度: 3
	////			切片的容量: 3
	//// slice2的内容:  [1 2 3]
	////			内存地址:  0xc0420024c0
	////			指向的内存地址: 0xc04200c360
	////			切片的长度: 3
	////			切片的容量: 3


	slice3 := []int{1,2,3}
	fmt.Println("slice3的内容: ",slice3)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice3))
	fmt.Printf("			指向的内存地址: %p\n",slice3)
	fmt.Printf("			切片的长度: %d\n", len(slice3))
	fmt.Printf("			切片的容量: %d\n",cap(slice3))

	slice4:= make([]int, len(slice3))
	copy(slice4,slice3)

	fmt.Println("slice4的内容: ",slice4)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice4))
	fmt.Printf("			指向的内存地址: %p\n",slice4)
	fmt.Printf("			切片的长度: %d\n", len(slice4))
	fmt.Printf("			切片的容量: %d\n",cap(slice4))

	// 输出
	// slice3的内容:  [1 2 3]
	//			内存地址:  0xc042002440
	//			指向的内存地址: 0xc04200c360
	//			切片的长度: 3
	//			切片的容量: 3
	// slice4的内容:  [1 2 3]
	//			内存地址:  0xc0420024c0
	//			指向的内存地址: 0xc04200c3a0
	//			切片的长度: 3
	//			切片的容量: 3
	// copy 赋值的话，那slice2指向的地址不等于slice2指向的地址
}