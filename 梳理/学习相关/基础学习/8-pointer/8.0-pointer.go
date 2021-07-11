package main

import "fmt"

func main(){
	var a = 10
	// 定义指针
	var p *int;
	fmt.Printf("p未赋值时:%x\n",p);
	// 判断空指针
	if p == nil{
		fmt.Printf("p现在是空指针\n")
	}

	p = &a;
	fmt.Printf("a的值是%d, a的地址是:%x\n",*p,p);
}

// 指针操作的& *和C一样