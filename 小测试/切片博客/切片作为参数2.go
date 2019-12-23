package main

import (
	"fmt"
)

func change(slice []int){
	slice[0] = 5
	fmt.Println("赋值后，函数内slice的内容: ",slice)
}


func main() {
	slice := []int{1,2,3}
	fmt.Println("函数外slice的内容: ",slice)
	change(slice)
	fmt.Println("函数执行后，函数外slice的内容: ",slice)
	// 输出
	// 函数外slice的内容:  [1 2 3]
	// 赋值后，函数内slice的内容:  [5 2 3]
	// 函数执行后，函数外slice的内容:  [5 2 3]
}
