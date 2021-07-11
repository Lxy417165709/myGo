package main

import (
	"fmt"
	"runtime"
	"time"
)


func timeCost(start time.Time,ret *time.Duration) {
	*ret = time.Since(start)
}

func myFunction(n int) {
	for i := 0; i < n; i++ {
	}
}

func myFunctionRev(i int) {
	if i<=0{
		return
	}
	myFunctionRev(i-1)
}
func showExecTime(f func(i int),i int,ret *time.Duration) {
	defer timeCost(time.Now(),ret)
	f(i)
}






func main() {

	for i := 0; i < 5; i++ {
		var dur time.Duration
		showExecTime(myFunction,100000000,&dur)
		fmt.Println(dur)
		showExecTime(myFunctionRev,1000000,&dur)
		fmt.Println(dur)
		fmt.Println()
		runtime.GC()
	}
	// 72.8053ms
	// 1.0408841s
	//
	// 69.8128ms
	// 79.7894ms
	//
	// 68.8138ms
	// 77.7912ms
	//
	// 67.8194ms
	// 80.7848ms
	//
	// 68.8149ms
	// 79.7826ms

	// 这是一个检查函数执行时间的程序~

	// 总结
	//		1. 首次递归很慢，之后的递归会快15倍左右。  (猜测：可能是第一次涉及到空间分配，之后空间被保留了)
	// 		2. 递归速度一般慢于迭代，迭代的速度是递归速度的10倍以上~

}
