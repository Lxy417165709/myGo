package main

import "fmt"

type Obj struct {

}
func (o *Obj)Do() {
	if o==nil{
		fmt.Println("不执行")
		return
	}
	fmt.Println("执行")
}


func main() {
	var o  *Obj
	fmt.Println(o)
	o.Do()

	// <nil>
	// 不执行
}

/*
总结：
	nil对象也能执行其方法！

*/
