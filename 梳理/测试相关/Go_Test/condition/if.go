package main

import "fmt"

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		fmt.Println(a, b, c)
	}
	// 1 2 3

	// 总结： 这题考查隐式代码块
}
