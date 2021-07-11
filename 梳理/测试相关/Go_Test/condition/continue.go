package main

import "fmt"

func main() {
RowLoop:
	for i := 0; i < 5; i++ {
		for t := 0; t < 5; t++ {
			fmt.Println(i, t)
			continue RowLoop
		}
	}
	// 0 0
	// 1 0
	// 2 0
	// 3 0
	// 4 0

	// 总结：continue 会跳过当前任务，继续执行标签处的循环
}
