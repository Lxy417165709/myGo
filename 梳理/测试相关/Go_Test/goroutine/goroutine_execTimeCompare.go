package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Function(i int) int {
	if i <= 0 {
		return 0
	}
	return i + Function(i-1)
}
func CountTime(fun func(int) int, num int, repeat int) time.Duration {
	before := time.Now()
	for i := 0; i < repeat; i++ {
		fun(num)
	}
	return time.Since(before)
}
func CountTimeGoroutine(fun func(int) int, num int, repeat int) time.Duration {
	before := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < repeat; i++ {
		wg.Add(1)
		go func() {

			fun(num)
			wg.Done()
		}()
	}
	wg.Wait()
	return time.Since(before)
}


func main() {
	for i:=1;i<=10;i++{
		runtime.GOMAXPROCS(i)
		fmt.Println(i,"N:",CountTime(Function, 1000, 10))
		fmt.Println(i,"G:",CountTimeGoroutine(Function, 1000, 10))
		fmt.Println()
	}

	// 结论：
	//       1. 在一定范围内，随着核数的增加，协程的并发效率会提升，之后会趋于平衡。
	//		 2. 可能由于协程创建的花费、扩栈、以及wg的花费，函数没并发执行时比并发执行时要快。


}
