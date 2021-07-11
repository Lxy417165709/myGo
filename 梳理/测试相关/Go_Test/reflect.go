package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	Name string
}

func main() {
	fmt.Println(reflect.ValueOf(Person{}).CanAddr())
	fmt.Println(reflect.ValueOf(&Person{}).CanAddr())
	//fmt.Println(reflect.ValueOf(Person{}).Elem().CanAddr())	// 报错
	fmt.Println(reflect.ValueOf(&Person{}).Elem().CanAddr())
	// false
	// false
	// true

	type IntAlias = int
	type MyInt int
	type MyMyInt MyInt

	var a int
	var b IntAlias
	var c MyInt
	var d MyMyInt

	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b).Kind())
	fmt.Println(reflect.TypeOf(c), reflect.ValueOf(c).Kind())
	fmt.Println(reflect.TypeOf(d), reflect.ValueOf(d).Kind())
	// int int
	// int int
	// main.MyInt int
	// main.MyMyInt int

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind())
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(c).Kind())
	fmt.Println(reflect.TypeOf(d), reflect.TypeOf(d).Kind())
	// int int
	// int int
	// main.MyInt int
	// main.MyMyInt int

	fmt.Println(reflect.TypeOf(&a), reflect.TypeOf(&a).Kind())
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&a)), reflect.TypeOf(unsafe.Pointer(&a)).Kind())
	fmt.Println(reflect.TypeOf(uintptr(unsafe.Pointer(&a))), reflect.TypeOf(uintptr(unsafe.Pointer(&a))).Kind())
	// fmt.Println(reflect.TypeOf(int(unsafe.Pointer(&a))))	 // 报错

	// *int ptr
	// unsafe.Pointer unsafe.Pointer
	// uintptr uintptr

	// 总结
	// canAddr() 需要 .Elem() 的加持，Elem() 作用在结构体类型会报错
	// Kind()返回的是基础类型，不管建立在什么类型上~
}
