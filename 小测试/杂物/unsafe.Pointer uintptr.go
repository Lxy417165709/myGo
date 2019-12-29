package 杂物

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "myString"
	p := &str
	p1 := unsafe.Pointer(p)
	//p2 := uintptr(p)	// 这样会报错
	p2 := uintptr(p1)

	fmt.Println("p:",p,reflect.TypeOf(p),unsafe.Sizeof(p))
	fmt.Println("p1:",p1,reflect.TypeOf(p1),unsafe.Sizeof(p1))
	fmt.Println("p2:",p2,reflect.TypeOf(p2),unsafe.Sizeof(p2))

	/*
		p: 0xc04204a1c0 *string 8
		p1: 0xc04204a1c0 unsafe.Pointer 8
		p2: 825741320640 uintptr 8
	*/
}
