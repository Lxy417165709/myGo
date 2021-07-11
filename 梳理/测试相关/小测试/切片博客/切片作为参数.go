package main

import (
	"fmt"
	"unsafe"
)

func show(slice []int){
	fmt.Println("函数内slice的内容: ",slice)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("			指向的内存地址: %p\n",slice)
	fmt.Printf("			切片的长度: %d\n", len(slice))
	fmt.Printf("			切片的容量: %d\n",cap(slice))
}


func main() {
	slice := []int{1,2,3}
	fmt.Println("函数外slice的内容: ",slice)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("			指向的内存地址: %p\n",slice)
	fmt.Printf("			切片的长度: %d\n", len(slice))
	fmt.Printf("			切片的容量: %d\n",cap(slice))

	show(slice)

	// 输出
	// 函数外slice的内容:  [1 2 3]
	//			内存地址:  0xc04204e3e0
	//			指向的内存地址: 0xc0420520a0
	//			切片的长度: 3
	//			切片的容量: 3
	// 函数内slice的内容:  [1 2 3]
	//			内存地址:  0xc04204e460
	//			指向的内存地址: 0xc0420520a0
	//			切片的长度: 3
	//			切片的容量: 3
	// 这种情况类似于使用等号赋值，函数形参和实参是两个不同的变量，所以有不同的内存地址，
	// 但是他们的值(对于切片来说就是指向地址，长度以及容量)是一样的
}
