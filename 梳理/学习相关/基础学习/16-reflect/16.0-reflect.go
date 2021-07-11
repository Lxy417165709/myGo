package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	age  int
}

func main() {
	var a = Student{"李學悅", 22}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	//main.Student
	//{李學悅 22}

	v := reflect.ValueOf(a)
	fmt.Println(v.FieldByName("我亂寫的"))
	fmt.Println(v.FieldByName("name"))
	fmt.Println(v.FieldByIndex([]int{0:1}))
	//<invalid reflect.Value>
	//李學悅

}

/*
	总结
	1. 通过反射,我们可以实现对结构体数据成员的动态访问,类似python的 myDict[<数据成员名>]
*/
