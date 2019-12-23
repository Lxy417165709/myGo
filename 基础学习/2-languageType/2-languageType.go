package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		a          = true
		b          = 10		// b的类型是int(可以是int32或int64),如果系统是32位,那就是int32,如果是64位,那就是int64
		c    int8  = 100
		d    int16 = 1000
		e    int32 = 10000
		f    int64 = 100000
		name       = 20500
	)
	fmt.Printf("%t %d %d %d %d %d\n", a, b, c, d, e, f)
	fmt.Printf("%c\n", name)
	// true 10 100 1000 10000 100000
	// 倔

	fmt.Printf("%d\n", int64(e)+f) // e + f 无法直接运算(go语言不会自动转换类型)
	// 110000

	// 查看变量占用字节
	fmt.Println(unsafe.Sizeof(b))
	// 8

}

/*
	总结
	1. go里面有int类型(8,16,32,64位)
	2. go里面有float类型(32,64位)
	3. go里面没有char类型
	4. go没有自动类型转换(int64不可以和int32直接运算)
	5. bool不能和其他类型相互转换
 */
