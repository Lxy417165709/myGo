package main

import "fmt"

// 正确输出: 10
func test1() int {
	x := 10
	defer func() {
		x++
	}()
	return x
}

// 正确输出: 10
func test2() (a int) {
	x := 10
	defer func() {
		x++
	}()
	return x
}

// 正确输出: 11
func test3() (x int) {
	x = 10 // x := 10 会报错
	defer func() {
		x++
	}()
	return x
}

// 正确输出: 11
func test4() (a int) {
	x := &a
	*x = 10
	defer func() {
		*x++
	}()
	return *x
}
func main() {
	fmt.Println("测试1:", test1())
	fmt.Println("测试2:", test2())
	fmt.Println("测试3:", test3())
	fmt.Println("测试4:", test4())
}
