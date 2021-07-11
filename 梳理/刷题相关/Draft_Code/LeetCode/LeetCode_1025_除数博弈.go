package main

import "fmt"

func divisorGame(N int) bool {
	return N%2==0
}


func main() {
	fmt.Println(divisorGame(2))
}

/*
	总结
	1. 这题博弈涉及到了一个知识点~~奇数的因子都是奇数~

*/
