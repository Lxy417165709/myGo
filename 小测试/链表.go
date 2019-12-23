package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Val int
	Next *Node
}

func main() {
	a := Node{1,nil}
	b := a
	c := &a
	d := new(Node)
	a.Val = 5
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&c),unsafe.Pointer(c))
	fmt.Println(unsafe.Pointer(&d),unsafe.Pointer(d))
	// 0xc0420481c0
	// 0xc0420481d0
	// 0xc042074018 0xc0420481c0
	// 0xc042074020 0xc0420481e0
}
/*
	总结
	1. 结构体的传递并非传指针，b:=a 并不是把a的值(指针)传递给b,而是再创建一个结构体存放a的信息，把该结构体的指针赋予b
	2. 可以参考值传递
	3. unsafe.Pointer(&c) 表示以指针的形式输出&c
	4. a := Node{1,nil} 则a的值不是指针，而是{1,nil}
	5. d := new(Node)	则d的值是指针，指向一个{x,x}
*/