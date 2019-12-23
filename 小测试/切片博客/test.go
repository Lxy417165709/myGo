package main

import "fmt"

func main() {
	mp2 := make(map[[2]int]int)		// 创建一个[2]int类型到int类型的map
	mp2[[2]int{1,2}] = 5			// 创建键值对
	mp2[[2]int{1,3}] = 5			// 创建键值对
	for k,v := range mp2{
		fmt.Println(k, v)
	}								// 遍历
	delete(mp2, [2]int{1, 2})		// 删除键值对
	for k,v := range mp2{
		fmt.Println(k, v)
	}								// 遍历
}
