
package main

import "fmt"

func main() {
	mp := make(map[int]int)
	for i:=0;i<10;i++{
		mp[i]=i
	}
	for k,v := range mp {
		fmt.Println(k,v)
	}
}

// 第一次运行
// 0 0
// 7 7
// 9 9
// 4 4
// 5 5
// 6 6
// 8 8
// 1 1
// 2 2
// 3 3

// 第二次运行
// 0 0
// 1 1
// 4 4
// 6 6
// 7 7
// 8 8
// 2 2
// 3 3
// 5 5
// 9 9

// 总结
// map 遍历的顺序是随机的！
