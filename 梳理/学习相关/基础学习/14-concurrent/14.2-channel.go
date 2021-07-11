package main

import "fmt"

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)	// 只是为了阻塞main协程,实现同步
//
//	go func() {
//		a := <-ch1
//		fmt.Println(a,"被提出")
//		ch2<-1
//	}()
//	time.Sleep(1e9)
//	go func() {
//		a:=10
//		fmt.Println(a,"丢进管道了")
//		ch1<-a
//		ch2<-2
//	}()
//	<-ch2
//	<-ch2
//	fmt.Println("程序结束")
//
//
//	// 10 丢进管道了
//	// 10 被提出
//	// 程序结束
//
//}

func main() {
	ch1 := make(chan int,10)
	for i:=0;i<10 ;i++  {
		ch1<-i
		if i==5{
			close(ch1)
			for x := range ch1{
				fmt.Println(x)

			}
		}
	}
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// panic: send on closed channel (出错了)
}

/*
	总结
	1. channel是安全的,多个goroutine(协程)同时操作的时候,同一时间只能有
	   一个goroutine存取数据
	2. channel达到容量后,无法进行存入,只能进行取出(存入操作会被阻塞)
       channel里面没数据时,只能进行存入,不能进行取出(取出操作会被阻塞)
	3. channel关闭后只能读 不能写

*/