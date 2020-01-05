package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	slice与unsafe.Pointer相互转换

	参考: https://www.w3cschool.cn/go_internals/go_internals-cyz6282h.html
*/

func main() {
	// slice -> unsafe.Pointer
	slice1 := make([]byte, 200)
	ptr := unsafe.Pointer(&slice1[0])
	fmt.Println("ptr:", ptr, reflect.TypeOf(ptr))

	// unsafe.Pointer -> slice (链接还提供了一些另外的构造方法)
	slice2 := ((*[1 << 10]byte)(ptr))[:500] // ptr指向一个长度为1024的数组
	fmt.Println("ptr:", ptr, reflect.TypeOf(ptr))
	fmt.Println("slice2:", len(slice2), cap(slice2))

	fmt.Println(&slice1[0],&slice2[0],slice1[0],slice2[0])
	slice1[0] = 100
	fmt.Println(&slice1[0],&slice2[0],slice1[0],slice2[0])
	// ptr: 0xc042084000 unsafe.Pointer
	// ptr: 0xc042084000 unsafe.Pointer
	// slice2: 200 1024
	// 0xc042072000 0xc042072000 0 0
	// 0xc042072000 0xc042072000 100 100

	// 居然指向了同一地址！！！
	// 应该是 [1 << 10]byte 并没有开辟新的数组，只是由于表示类型


	var ptr1 unsafe.Pointer
	var s1 = struct {
		addr uintptr
		len int
		cap int
	}{uintptr(ptr1), 500, 500}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	slice1 = s
	fmt.Println(reflect.TypeOf(s),reflect.TypeOf(slice1))
	// []uint8 []uint8


	ss := []int{}
	fmt.Println(ss,len(ss))
	// [] 0
}
