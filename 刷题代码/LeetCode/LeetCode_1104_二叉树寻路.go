package main

import (
	"fmt"
	"math"
)

func pathInZigZagTree(label int) []int {
	ans = make([]int,0)
	solve(label)
	for i:=0;i< len(ans)/2;i++  {
		ans[i],ans[len(ans)-1-i] = ans[len(ans)-1-i],ans[i]
	}
	return ans
}

var ans []int

func solve(label int){
	ans = append(ans, label)
	if label ==1 {
		return
	}
	base:=int(math.Pow(2,float64(int(math.Log2(float64(label))))))
	x := (label -base)/2
	y:=base - x-1
	solve(y)
}

func main() {
	fmt.Println(pathInZigZagTree(7))
}
/*
	总结
	1. 这题我是通过找规律后递归AC的 0.0.
			1  		1
			2  		1 2
			3		1 3
			4		1 3 4
			5		1 3 5
			6		1 2 6
			7		1 2 7
			8		1 2 7 8
			9		1 2 7 9
			10		1 2 6 10
			11		1 2 6 11
	2. 官方有用位运算做的 0.0.
*/