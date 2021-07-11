package main

import "fmt"

type Writer interface {
	Write()
}
type Author struct{
	Writer
}
type Other struct{

}
func (a Author)Write(){
	fmt.Println("Author")
}
func (o Other)Write(){
	fmt.Println("Other")
}

type Buthor struct {
	Writer
}
func main() {
	ao := Author{Other{}}
	ao.Write()
	a := Author{}
	a.Write()

	// b := Buthor{}
	// b.Write()		// 编译不报错，运行时报错 invalid memory address or nil pointer dereference



	// Author
	// Author

	// 总结
	//    1. 结构体方法访问优先级 > 嵌入接口方法优先级
}
