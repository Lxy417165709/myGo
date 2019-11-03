package main

import (
	"fmt"
	"sync"
)

//func main() {
//	m := make(map[int]int)
//	var wg sync.WaitGroup	// 阻塞main协程
//	wg.Add(10)
//	for i:=0;i<10 ;i++  {
//		go func() {
//			m[i]=i			// 写
//			fmt.Println(m)	// 读
//			wg.Done()
//		}()
//
//	}
//	wg.Wait()
//	fmt.Println("程序结束")
//	fmt.Println(m)	// 读
//
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[10:10 3:5]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// map[3:5 10:10]
//	// 程序结束
//	// map[3:5 10:10]
//	// (这很诡异...)
//}

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup	// 阻塞main协程
	wg.Add(10)
	for i:=0;i<10 ;i++  {
		go func(j int) {
			m[j]=j			// 写
			fmt.Println(m)	// 读
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println("程序结束")
	fmt.Println(m)
	// map[0:0 9:9 4:4]
	// map[0:0 9:9 4:4 6:6 1:1]
	// map[0:0 9:9 4:4]
	// map[0:0 9:9 4:4 6:6 1:1 7:7]
	// map[0:0 9:9 4:4 6:6 1:1 7:7 8:8]
	// map[9:9 4:4 6:6 1:1 7:7 8:8 5:5 0:0]
	// map[6:6 1:1 7:7 8:8 2:2 3:3 9:9 4:4 0:0 5:5]
	// map[0:0 5:5 1:1 7:7 8:8 2:2 9:9 4:4 6:6]
	// map[0:0 9:9 4:4]
	// map[0:0 9:9 4:4 6:6]
	// 程序结束
	// map[9:9 1:1 6:6 7:7 8:8 0:0 2:2 3:3 4:4 5:5]
}