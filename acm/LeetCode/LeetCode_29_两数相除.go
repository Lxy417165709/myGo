package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	if dividend==-(1<<31) && divisor==-1 {
		return 1<<31-1
	}
	flag:=0

	if dividend<0 {
		if divisor>0 {
			flag = 1
		}
	}else{
		if divisor<0 {
			flag = 1
		}
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	sli := make([]int,0)
	weight := make([]int,0)

	w:=1

	for ;divisor<=dividend ; {
		sli = append(sli,divisor)
		weight = append(weight,w)
		w = w + w
		divisor = divisor + divisor
	}
	ans:=0
	for i:= len(sli)-1;i>=0 ;i--  {
		if sli[i]<=dividend {
			dividend = dividend - sli[i]
			ans+=weight[i]
		}
	}

	//fmt.Println(sli,weight)
	if flag==1 {
		return -ans
	}else{
		return ans
	}
}

func main() {
	fmt.Println(divide(10,3))
}

/*
	总结
	1. 该题目还是用了二分思维
	2. 注意溢出判断
*/