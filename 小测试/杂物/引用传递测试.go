package 杂物

import "fmt"

type Student struct {
	Name string
}

func main() {
	a:= new(Student)
	b:=a
	fmt.Println(a,b)
	a.Name = "我我我"
	fmt.Println(a,b)
}

/*
	总结
	1. 结构体赋值传递的是指针，两个指针指向同一对象，所以通过a的改变会影响b

*/

