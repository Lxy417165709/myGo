package main

import "fmt"

type Student struct {
	Name string
}
func (s Student)say(){
	fmt.Println("我叫" + s.Name)
	s.Name = "哈哈哈"

	fmt.Printf("函数内变量的地址: %x\n", &s  )
	fmt.Printf("函数内变量指向的地址: %x\n", s  )
}


func main() {
	s:=Student{"李学悦"}
	s.say()
	fmt.Printf("变量的地址: %x\n", &s  )
	fmt.Printf("变量指向的地址: %x\n", s  )

	// 我叫李学悦
	// 我叫李学悦
	// 函数内变量的地址: &{e59388e59388e59388}
	// 函数内变量指向的地址: {e59388e59388e59388}
	// 变量的地址: &{e69d8ee5ada6e682a6}
	// 变量指向的地址: {e69d8ee5ada6e682a6}
}

/*
	总结
	1. func (s *Student)say() 此时传递的是指向对象的指针，对s指向的对象的修改会影响传入对象(注意go语言*&有时候可以省略不写)
	2. func (s Student)say()  此时传递的是对象的信息，对s指向的对象的修改不会影响传入对象
*/