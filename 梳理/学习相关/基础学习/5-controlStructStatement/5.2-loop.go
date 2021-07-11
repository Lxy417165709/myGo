package main

import "fmt"

func main() {

	// 1. for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i : %d\n", i)
	}
	//for true {
	//	fmt.Println("这是一个无限循环")
	//}

	// 2. 控制语句
	// 2.0 continue (不执行下面的代码,返回继续循环)
	for i := 0; i < 3; i++ {
		for t := 0; t < 3; t++ {
			if t >= 1 {
				continue
			}
			fmt.Printf("%d ", t)
		}
		fmt.Println()
	}
	// 0
	// 0
	// 0

	// 2.1 break (跳出循环)
X:
	for i := 0; i < 3; i++ {
		for t := 0; t < 3; t++ {
			if t >= 1 {
				break X
			}
			fmt.Printf("%d ", t)
		}
		fmt.Println()
	}
	// 0

	// 2.2 goto	(到指定处执行)
Y:
	for i := 0; i < 3; i++ {
		for t := 0; t < 3; t++ {
			if t >= 1 {
				goto Y
			}
			fmt.Printf("%d ", t)
		}
		fmt.Println()
	}
	// 0 0 0 0 0 ....... (无限)


	// 2.3 循环遍历数组
	array := [...]int{1,2,3,4,5}
	for index,value := range array {
		fmt.Println(index,value)
	}

	//0 1
	//1 2
	//2 3
	//3 4
	//4 5
}

/*
	总结
	1. go语言只有for循环,没有while和do...while循环
	2. go语言的continue,goto和C++语言一样,但是break不一样,go语言的break可以指定需要跳出哪层循环
*/
