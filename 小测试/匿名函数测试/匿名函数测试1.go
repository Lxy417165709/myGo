package main

import "fmt"

func main() {

	var foo func(int) int
	// 匿名函数
	foo = func(n int) int {
		if n == 0 {
			return 0
		}
		return foo(n-1) + n
	}
	fmt.Println(foo(10))
	// 55
}
