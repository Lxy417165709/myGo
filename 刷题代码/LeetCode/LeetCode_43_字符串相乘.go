package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	s := [500]int{}
	if len(num1)<len(num2) {
		num1,num2 = num2,num1
	}
	zero := len(num1) - len(num2)
	for i:=1;i<= zero;i++  {
		num2 = "0" + num2

	}

	num1 = rev(num1)
	num2 = rev(num2)
	length := len(num1)


	for i:=0;i<=length-1 ;i++  {
		sum := 0
		u:=0
		for t:=0;t<=length-1 ;t++  {
			sum = s[i+t] + int((num1[i]-'0') * (num2[t]-'0')) + sum
			s[i+t] = sum%10
			sum = sum /10
			u = i + t

		}
		for ;sum!=0 ;  {
			u++
			sum = s[u] + sum
			s[u] = sum%10
			sum = sum /10
		}

	}


	ans:=""
	for i:=2*length;i>=0 ;i--  {
		if s[i]==0 {
			continue
		}
		for t:=i;t>=0 ;t--  {
			ans =  ans + string(s[t] +  '0')
		}
		break
	}
	if ans=="" {
		ans="0"
	}
	return ans
}

func rev(s string) string{

	ans:=""
	for i:=0;i<len(s) ;i++  {
		ans = ans + string(s[len(s)-1-i])
	}

	return ans
}

/*
	总结
	1. 这题目模拟乘法，虽然写出来了，但是还可以更简洁的，比如在求结果位数有多少的时候
	2. 可以不必添0就能计算了
	3. 进位循环可以改进，可以不需要那个for循环

*/


func main() {
	fmt.Println(multiply("2","3"))
}
