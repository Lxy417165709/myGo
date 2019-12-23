package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := new(int) // 此时a是指针变量,值是一个内存地址
	fmt.Println(&a, a, *a,unsafe.Sizeof(a))

	var b *int
	fmt.Println(&b, b,unsafe.Sizeof(b))

	// 0xc042074018 0xc042054080 0 8
	// 0xc042074028 <nil> 8
}

/*
	总结
	1. go语言的指针操作,new,*,&和C语言一样
*/
