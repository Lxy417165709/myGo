package 杂物

import "fmt"

func change(sl []int)  {
	sl = append(sl,1)
	sl[0]=1
	fmt.Println("函数执行后:",sl)
}

func main() {
	var sl []int
	sl = make([]int,1)
	fmt.Println("刚开始:",sl)
	change(sl)
	fmt.Println("最后:",sl)
}

/*
	总结
	1. append会改变指针 0.0..
*/