package main

import (
	"fmt"
	"time"
)

func produce(c chan int){
	for i:=1;i<=100 ;i++  {
		c <- i
		fmt.Println(i,"被生产出来了~","现在C有",len(c),"物品")
		time.Sleep(200*time.Millisecond)
	}
}
func consume(c chan int){

	for i:=1;i<=100 ;i++  {
		k := <-c
		fmt.Println(k,"被消耗咯~")
		time.Sleep(300*time.Millisecond)
	}
}


func main(){
	var c = make(chan int,50)
	go produce(c)
	go consume(c)

	time.Sleep(1000000000000*time.Millisecond)	// 阻塞main协程
}


// channel在到达容量后,会阻塞线程
// 不指定容量时(容量为0),当channel有数据,则会阻塞指导到数据被提出