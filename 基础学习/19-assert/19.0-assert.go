package main

import "fmt"

func main() {
	var a interface{} = 1234

	fmt.Printf("%T\n",a)
	// int

	b:=3.4
	fmt.Printf("%T\n",b)
	// float64

	_,ok := a.(float64)
	fmt.Println(ok)
	// false

}

/*
	总结
	1.  fmt.Printf("%T\n",<变量>)  这个可以获取变量类型
	2.  var a interface{}
		result,ok := a.(type)
		如果a的类型是type,那result == a, ok == true
		如果a的类型是type,那result == 默认值, ok == false

*/