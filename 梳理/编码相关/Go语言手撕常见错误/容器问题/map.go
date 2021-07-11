package main

import (
	"fmt"
)

// 目的: 定义一个map并创建一个键值对
func problemOne(){
	var mp map[int]int
	mp[1] = 5
	fmt.Println(mp[1])

	// panic: assignment to entry in nil map
	// 出现问题: map未分配空间
	// 解决方法: var mp map[int]int    ---->    mp := map[int]int{}
}

func main() {
	problemOne()
}