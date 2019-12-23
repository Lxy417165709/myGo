package _1nod

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	if n <= 2 {
		fmt.Println(n)
		return
	}
	var a int64 = n*(n-1)*(n-2)
	var b int64 = n*(n-1)*(n-3)
	var c int64 = (n-1)*(n-2)*(n-3)
	if n%2==0{
		if n%3==0{
			fmt.Println(c)
		}else{
			fmt.Println(b)
		}
	}else{
		fmt.Println(a)
	}





}

/*
	总结
	1. 这题目要分类思维...还要会一些数论

*/