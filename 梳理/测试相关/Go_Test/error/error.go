package main

import "fmt"

// 返回值被屏蔽
func Foo() (err error) {
	if err := Bar(); err != nil {
		return
	}
	return
}
func Bar() error{
	return fmt.Errorf("自定义错误")
}

func main() {
	fmt.Println(Foo())
	// 出现错误：..\..\..\..\Desktop\Go_Test\error.go:8:3: err is shadowed during return
}
