package main

import (
	"fmt"
	"time"
)

func say(sentence string){
	for i:=0;i<5 ;i++  {
		time.Sleep(100*time.Millisecond)	// 让当前线程休息一下 也可以用  runtime.Gosched()
		//runtime.Gosched()
		fmt.Println(sentence)
	}
}

func main(){
	/*
		go say("aaa")
		go say("bbb")
		这样的话什么都没输出...
	*/
	go say("aaa")
	go say("bbb")

	// 以下是为了防止main函数结束
	/*
	方式1
	var str string
	fmt.Scan(&str)

	*/
	// 方式2
	time.Sleep(2000*time.Millisecond)
}

// 线程调度要时间...,如果main协程结束了,那其他协程都会结束(不会执行了)