package main

import "fmt"

func main(){
	a,b := 10,20
	// 算数运算符 + - * / % ++ --
	fmt.Println(a + b)

	// 关系运算符 > < = >= <= == !=
	fmt.Println((a == b))

	// 逻辑运算符 && || !
	c,d := false, true
	fmt.Println(c && d)
	fmt.Println(!d)

	// 位运算符(对整数在内存中的二进制位进行操作)	& | ^ << >>
	e, f := 5,2
	fmt.Println(e & f)
	fmt.Println(^-1)	// 单目时是取反，双目时是异或
	fmt.Println(f & 1)	// 相当于 f % 2

	// 赋值运算符 = += -= *= /=  %= &= |= ~= <<= >>=
	f <<= 3				// 相当于 f * (2的3次方)
	fmt.Println(f)
}
