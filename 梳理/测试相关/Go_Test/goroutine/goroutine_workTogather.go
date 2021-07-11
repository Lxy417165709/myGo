package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
