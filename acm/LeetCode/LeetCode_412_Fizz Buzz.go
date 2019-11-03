package main

import (
	"fmt"
	"strconv"
)
// 暴力写法
//func fizzBuzz(n int) []string {
//	slice:=make([]string,0)
//	for i:=1;i<=n ;i++  {
//		if i%3==0 {
//			if i%5==0 {
//				slice  = append(slice, "FizzBuzz")
//			}else{
//				slice  = append(slice, "Fizz")
//			}
//		}else{
//			if i%5==0 {
//				slice  = append(slice, "Buzz")
//			}else{
//				slice  = append(slice, strconv.Itoa(i))
//			}
//		}
//	}
//	return slice
//}
// 优雅写法
func fizzBuzz(n int) []string {
	slice:=make([]string,0)
	for i:=1;i<=n ;i++  {
		s:=""
		if i%3==0 {
			s+="Fizz"
		}
		if i%5==0 {
			s+="Buzz"
		}
		if s=="" {
			s+=strconv.Itoa(i)
		}
		slice = append(slice, s)

	}
	return slice
}
func main() {
	fmt.Println(fizzBuzz(15))
}
/*
	总结
	1. 这是一道超级简单的题目
	2. 官方有更优雅的写法 0.0.。所以我也写了一次
*/