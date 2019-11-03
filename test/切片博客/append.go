package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := make([]int,2,3)	// 另外一种创建切片的方式
	fmt.Println("切片的内容: ",slice)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("			指向的内存地址: %p\n",slice)
	fmt.Printf("			切片的长度: %d\n", len(slice))
	fmt.Printf("			切片的容量: %d\n",cap(slice))

	slice = append(slice,5)
	fmt.Println("追加一个元素5后")
	fmt.Println("切片的内容: ",slice)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("			指向的内存地址: %p\n",slice)
	fmt.Printf("			切片的长度: %d\n", len(slice))
	fmt.Printf("			切片的容量: %d\n",cap(slice))


	slice = append(slice,10)
	fmt.Println("再追加一个元素10后")
	fmt.Println("切片的内容: ",slice)
	fmt.Println("			内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("			指向的内存地址: %p\n",slice)
	fmt.Printf("			切片的长度: %d\n", len(slice))
	fmt.Printf("			切片的容量: %d\n",cap(slice))


	// 输出
	// 切片的内容:  [0 0]
	//			内存地址:  0xc04204e3e0
	//			指向的内存地址: 0xc0420520a0
	//			切片的长度: 2
	//			切片的容量: 3
	// 追加一个元素5后
	// 切片的内容:  [0 0 5]
	//			内存地址:  0xc04204e3e0
	//			指向的内存地址: 0xc0420520a0
	//			切片的长度: 3
	//			切片的容量: 3
	// 再追加一个元素10后
	// 切片的内容:  [0 0 5 10]
	//			内存地址:  0xc04204e3e0
	//			指向的内存地址: 0xc042068060
	//			切片的长度: 4
	//			切片的容量: 6
}

