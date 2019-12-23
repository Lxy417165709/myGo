package main

import (
	"fmt"
)

type Student struct {
	id int
	name string
	sex string
}

func main(){
	array := [...]int{1,2,3,4,5}
	for index,value := range array {
		fmt.Println(index,value)
	}

	//0 1
	//1 2
	//2 3
	//3 4
	//4 5

	/*
	student := Student{1,"李学悦","男"}
	for key,value := range student{
		fmt.Println(key,value)
	}
		结构体不能用range遍历
	*/

	m := map[string]string{"num1":"小李","num2":"小王","num3":"小白"}
	for key,value := range m{
		fmt.Println(key,value)
	}
	//num1 小李
	//num2 小王
	//num3 小白


	c := make(chan int,10)
	for i:=0;i<10 ;i++  {
		c<-i

	}
	for number := range c{
		println(number)
	}

	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	//fatal error: all goroutines are asleep - deadlock!


}
