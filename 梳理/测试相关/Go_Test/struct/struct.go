package main

import "fmt"

type A struct{
	int
}

func (a *A)Show(){
	fmt.Println("a")
}

type B struct{
	int
}

type A1 struct{
	a int
}
type B1 struct{
	a int
}

type C struct {
	A
	B
}

func (c *C)Show(){
	fmt.Println("c")
}

type C1 struct {
	A1
	B1
}

type D struct {
	A
	int
}

func main(){
	c := C{A{1},B{2}}
	//c1 := C1{A1{1},B1{5}}
	d := D{A{1},5}

	//fmt.Println(c.int)	// 报错 ---> 这会引发二义性
	//fmt.Println(c1.a)	    // 报错 ---> 这会引发二义性
	fmt.Println(d.int)	    // 没报错 ---> 5
	c.Show()				// 输出 c

	// 总结:
	//		1. Go编译器不知道怎么处理两个嵌入结构体的同名字段。这种情况编译器会报错。
	//      2. 本结构体的字段、方法访问优先级 > 嵌入字段、方法访问优先级

}
