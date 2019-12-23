package main

import "fmt"

var m [50]int

func tribonacci(n int) int {
	m[0] = 0
	m[1] = 1
	m[2] = 1
	for i:=3;i<=n ;i++  {
		m[i]= m[i-1] + m[i-2] + m[i-3]
	}
	return  m[n]


}


func main() {
	fmt.Println(tribonacci(2))
}


/*
	总结
	1. 这是斐波那契数的变形~
*/