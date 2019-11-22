package main

import "fmt"

func main() {
	mp:=make(map[[2]int]int)

	a:=[2]int{1,2}
	b:=[2]int{1,2}
	fmt.Println(mp[b])
	mp[a]=1
	fmt.Println(mp[b])
	// 0
	// 1


}

/*
	总结
	1. 以数组为Key的map,key的指标根据的是数组的内容
*/
