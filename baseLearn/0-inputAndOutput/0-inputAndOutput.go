package main

import (
	"fmt"
)

func main() {
	var a string
	//fmt.Println("请输入字符串")
	//
	//// Fscan
	//fmt.Fscanln(os.Stdin, &a)
	//fmt.Fscan(os.Stdin, &a)
	// 输入包括 Scan,Scanln,Scanf
	//fmt.Scanln(&a,&b)
	fmt.Printf("%s",a)

	// Sprint
	s := fmt.Sprintln("abc")	// 返回值是 abc\n
	fmt.Println(s)

	s = fmt.Sprintf("%d %d",1,2)	// 返回值是 1 2
	fmt.Println(s)

	s = fmt.Sprint("abc")	// 返回值是 abc
	fmt.Println(s)

	s = fmt.Sprintf("%04d %04d",1,2555)	// 返回值是 0001 25555
	fmt.Sscan(s,"123")
	fmt.Println(s)
}

/*
	总结:
	1. 输入 有F开头的表示输入到某个地方,例如 控制台
	2. 输出 有F开头的表示输出到某个地方,例如 控制台,网页...
	3. 输入函数 Scanf,Scanln,Scan    (前面可以加F,S)	(scanf,scanln是一樣的)
	4. 输出函数 Printf,Println,Print (前面可以加F,S)
	5. 没有F开头的输入输出,底层实现都是用F开头的输入输出
	6. 以S开头的scan...(还不清楚)
	7. 以S开头的print相当于把数据转换为字符串



 */
