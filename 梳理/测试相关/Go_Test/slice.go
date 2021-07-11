package main

import "fmt"

func main() {
	s := make([]*int,0)
	s = append(s,nil)

	// 总结
	// 		对于引用类型的切片，切片可以追加nil


	s1 := make([]int,5,10)
	//fmt.Println(s1[9])
	// panic: runtime error: index out of range [9] with length 5

	fmt.Println(s1[:9])
	// [0 0 0 0 0 0 0 0 0]

	// 总结：下标访问的索引最大为 长度 - 1， 而范围访问最大为 容量 - 1

}
