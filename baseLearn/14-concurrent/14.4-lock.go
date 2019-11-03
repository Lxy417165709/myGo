package main

import (
	"fmt"
	"sync"
	"time"
)
var (
	k int
	rwm sync.Mutex
)

// 没有加锁
func fun1(j int){

	for i:=0;i<10;i++{
		k=i
		time.Sleep(1e8)	// 暂停当前协程,在这个期间k可以被其他的协程修改,所以会出现 i != k 的情况
		fmt.Println(j,i,k)

	}
}
// 加锁后
func fun2(j int){

	for i:=0;i<10;i++{
		rwm.Lock()
		k=i
		time.Sleep(1e8)	//这个可以省略了
		fmt.Println(j,i,k)
		rwm.Unlock()
		// 加锁后,其他协程只能等待这个协程解锁后才能读入和写出(rwm是读写锁)
		// 所以一定会有 i == k

	}
}

//func main() {
//
//	for i:=0;i<3;i++{
//		go fun1(i)
//	}
//
//	time.Sleep(5e9)
//	fmt.Println("程序结束")
//
//
//	// 1 0 0
//	// 0 0 0
//	// 2 0 1
//	// 0 1 1
//	// 1 1 2
//	// 2 1 2
//	// ...
//	// 程序结束
//	// 为什么会这样呢?因为k在其他协程被修改了!!!
//}

func main() {

	for i:=0;i<3;i++{
		go fun2(i)
	}

	time.Sleep(5e9)
	fmt.Println("程序结束")


	// 1 7 7
	// 2 7 7
	// 0 9 9
	// 1 8 8
	// 2 8 8
	// 1 9 9
	// 2 9 9
	// ...
	// 程序结束
	// 这个是正常输出
}


/*
	总结
	1. fmt.Print(x)命令的执行需要时间,在这个执行的同时,其他协程也可以工作(不加锁的话,其他协程在这个期间可以修改x的值)
	2. 加锁后协程是安全的,但是会损失程序执行的效率(因为其他协程将会等待读写,直到该协程读写锁被解)
*/