package main

import (
	"fmt"
	"sync"
)

//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(10)
//	go func() {
//		wg.Done()
//	}()
//
//	wg.Wait()
//  fmt.Println("程序结束")
//	// fatal error: all goroutines are asleep - deadlock! (因为wg还有9个单元,wg.Wait()只有当wg里面没有单元后,他才会结束阻塞)
//
//}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	go func() {
		for i:=0;i<10 ;i++  {
			fmt.Println(i)
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println("程序结束")
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 程序结束

}

/*
	总结
	1. wg.Add, wg.Done, wg.Wait 应该配套使用,缺一不可 (保证有意义的情况下)
	2. wg.Add 向wg添加元素
	3. wg.Done 从wg取出一个元素
	4. wg.Wait 等待wg为空后,执行后面的代码 (wg不为空则会一直阻塞)
*/