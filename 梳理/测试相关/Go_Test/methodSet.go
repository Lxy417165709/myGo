package main

import (
	"fmt"
	"reflect"
)

type Person struct {
}

// 要可见的！！！
func (p Person) M1()  {}
func (p *Person) M2() {}

type Man interface {
	M1()
	M2()
}

func methodCount(itfc interface{}) int {
	tp := reflect.TypeOf(itfc)
	for i := 0; i < tp.NumMethod(); i++ {
		fmt.Println(tp.Method(i))
	}
	return tp.NumMethod()
}
func main() {
	p := Person{}
	fmt.Println(methodCount(p))  // 1 -> 说明 Person 只有一个方法，为 M1
	fmt.Println(methodCount(&p)) // 2 -> 说明 Person 只有两个方法，为 M1、M2

	// {M1  func(main.Person) <func(main.Person) Value> 0}
	// 1
	// {M1  func(*main.Person) <func(*main.Person) Value> 0}
	// {M2  func(*main.Person) <func(*main.Person) Value> 1}
	// 2

	// 以上说明：Person  只有一个方法，为 M1
	//			 *Person 有两个方法，  为 M1、M2

	var m Man
	//m = p		-> 报错，说明 Person 没有实现 Man接口
	m = &p //	-> 没报错，说明 *Person 实现了 Man接口

	_ = m
}
