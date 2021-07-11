package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	for v := range ch {
		fmt.Println(v)
	}
	// 1
	// 2
	// 3
	// 4
	// 5
	// fatal error: all goroutines are asleep - deadlock!

	// 总结：for range对于通道时，通道关闭后，for range才能结束。

	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	//1
	//2
	//3
	//4
	//5

	// 此时没有任何异常
}
