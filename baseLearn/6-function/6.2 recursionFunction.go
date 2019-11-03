package main

import "fmt"

func fbnq(n int) int {
	if n==0 || n==1{
		return n
	}else{
		return fbnq(n-1)+fbnq(n-2)
	}
}
func main(){
	fmt.Println(fbnq(10))
	for i:=0;i<10;i++{
		fmt.Printf("%d ",fbnq(i))
	}
	fmt.Println()
	// 0 1 1 2 3 5 8 13 21 34
}

// 递归和C语言一样
