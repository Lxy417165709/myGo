package main

import (
	"fmt"
	"sync/atomic"
	"time"
)
/*
func main()  {
	var money int64
	for i:=0;i<2000;i++{
		go func(i int) {
			money++
		}(i)
	}

	time.Sleep(2000)
	fmt.Println(money)
	// C:\Users\hasee\Desktop\Go_Test>go run -race race.go
	//==================
	//WARNING: DATA RACE
	//Read at 0x00c0000a2068 by goroutine 8:
	//  main.main.func1()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:12 +0x3f
	//
	//Previous write at 0x00c0000a2068 by goroutine 7:
	//  main.main.func1()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:12 +0x55
	//
	//Goroutine 8 (running) created at:
	//  main.main()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:11 +0x8f
	//
	//Goroutine 7 (finished) created at:
	//  main.main()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:11 +0x8f
	//==================
	//994

	// 这样会产生竞态问题
}
*/

func main()  {
	var money int64
	for i:=0;i<200000;i++{
		go func(i int) {
			atomic.AddInt64(&money,1)
		}(i)
	}

	time.Sleep(2000)
	fmt.Println(money)

	// C:\Users\hasee\Desktop\Go_Test>go run -race race.go
	//==================
	//WARNING: DATA RACE
	//Read at 0x00c0000100a0 by main goroutine:
	//  main.main()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:52 +0xc1
	//
	//Previous write at 0x00c0000100a0 by goroutine 7:
	//  sync/atomic.AddInt64()
	//      C:/Language/Go/src/runtime/race_amd64.s:276 +0xb
	//  main.main.func1()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:46 +0x48
	//
	//Goroutine 7 (finished) created at:
	//  main.main()
	//      C:/Users/hasee/Desktop/Go_Test/race.go:45 +0x8f
	//==================
	//200000
	//Found 1 data race(s)
	//exit status 66

	// 这个也会出现竞态条件...不知道为啥，不过结果运算都是正确的
}
