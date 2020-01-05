package main

import "fmt"


func main() {
	foo := make(map[int]int)
	foo[10] = 5
	foo[11] = 10
	fmt.Println(foo)
	for k, v := range foo {
		fmt.Println(k, v, &k, &v)
	}
	// 第一种情况
	// map[10:5 11:10]
	// 10 5 0xc0420560b8 0xc0420560d0
	// 11 10 0xc0420560b8 0xc0420560d0

	// 第二种情况
	// map[11:10 10:5]
	// 10 5 0xc0420560b8 0xc0420560d0
	// 11 10 0xc0420560b8 0xc0420560d0
}
