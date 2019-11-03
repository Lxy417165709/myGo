package main

import "fmt"

func main() {
	a := 1
	b := a
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	// 0xc042054080 0xc042054088
	// 1 1

	b = 2
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	// 0xc042054080 0xc042054088
	// 1 2

	c := 1.5
	p1:=&c
	p2:=p1
	fmt.Println(p1, p2,c)
	fmt.Println(&p1, &p2)
	// 0xc0420540d0 0xc0420540d0 1.5
	// 0xc042074020 0xc042074028

	*p2 = 5
	fmt.Println(p1, p2,c)
	fmt.Println(&p1, &p2)
	// 0xc0420540d0 0xc0420540d0 5
	// 0xc042074020 0xc042074028

	d := new(int)
	*d = 5
	e := d
	*e = 10
	fmt.Println(d,e,*d)
	// 0xc04206e0c8 0xc04206e0c8 10


}

/*
	总结
	1. 指针也需要空间来存储 &p表示取出这个指针的地址(其实和之前自己的理解一样)
	2. 地址-变量名-值,  &变量名 为取出变量的地址
						如果变量是指针,*变量名为 取出变量名的值
	3. go的new和c++语言的一样(除了写法)
*/