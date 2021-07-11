package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "hello world"
	fmt.Println(&str,str) // 0xc00002e1f0 不会报错
	fmt.Println((*rune)(unsafe.Pointer(&str)))
	p := (*rune)(unsafe.Pointer(&str))
	*p = 'h'	//  panic: runtime error: invalid memory address or nil pointer dereference

	// 总结
	// 1. 字符串可以取址，但其指向的内存不允许改变(因为它位于只读段)
}
