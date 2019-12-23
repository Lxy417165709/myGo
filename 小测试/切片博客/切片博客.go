package 切片博客

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [5]int{1,2,3,4,5}
	slice := array[1:]			// 与上面的代码相比，只有这做了更改


	fmt.Println("数组所在内存地址: ",unsafe.Pointer(&array))
	fmt.Println("切片所在内存地址: ",unsafe.Pointer(&slice))
	fmt.Printf("切片指向的内存地址: %p\n",slice)
	fmt.Printf("切片的长度: %d\n", len(slice))
	fmt.Printf("切片的容量: %d\n",cap(slice))
	// 数组所在内存地址:  0xc04200a330
	// 切片所在内存地址:  0xc042002440
	// 切片指向的内存地址: 0xc04200a338
	// 切片的长度: 4
	// 切片的容量: 4

	// 可以看出，切片指向的内存地址为数组所在内存地址(即数组元素0的内存地址)
}
