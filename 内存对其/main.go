package main

import (
	"fmt"
	"unsafe"
)

type Foo1 struct {
	a byte
}

type Foo2 struct {
	a byte
	b int32
}

type Foo3 struct {
	a int32
}

type Foo4 struct {
	a byte
	b int32
	c int64
}

type Foo5 struct {
	a byte
	c int64
	b int32
}


func main() {
	fmt.Println(unsafe.Sizeof(Foo1{}))
	fmt.Println(unsafe.Sizeof(Foo2{}))
	fmt.Println(unsafe.Sizeof(Foo3{}))
	fmt.Println(unsafe.Sizeof(Foo4{}))

	fmt.Println(unsafe.Alignof(Foo4{}.a)) // 1
	fmt.Println(unsafe.Offsetof(Foo4{}.a)) // 0
	fmt.Println(unsafe.Alignof(Foo4{}.b)) // 4
	fmt.Println(unsafe.Offsetof(Foo4{}.b)) // 4
	fmt.Println(unsafe.Alignof(Foo4{}.c)) // 8
	fmt.Println(unsafe.Offsetof(Foo4{}.c)) // 8

	fmt.Println(unsafe.Alignof(Foo4{})) // 8
	//fmt.Println(unsafe.Offsetof(Foo4{})) // invalid expression unsafe.Offsetof(Foo4 literal)

	fmt.Println(unsafe.Sizeof(Foo5{}))	// 24
	fmt.Println(unsafe.Alignof(Foo4{})) // 8
}
