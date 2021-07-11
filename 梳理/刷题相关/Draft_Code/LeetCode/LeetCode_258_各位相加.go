package main

import "fmt"

// 打表找规律
func addDigits(num int) int {
	ans:=num
	for ;ans>=10 ;  {
		y:=0
		for ;num!=0 ;  {
			y+=num%10
			num=num/10
		}
		ans = y
		num = y
	}
	return ans
}
func addDigits1(num int) int {
	if num==0 {
		return 0
	}
	ans:=num%9
	if ans==0{
		ans = 9
	}
	return ans
}
func main() {

	for i:=0;i<=0 ;i++  {
		fmt.Println(i,addDigits(i),addDigits1(i))
	}
}
/*
	总结
	1. 这题刚开始知道怎么做，但是要求O(1)时间复杂度，于是打表找规律做出来了0.0..
*/