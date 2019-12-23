package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	x:=solve(nums)
	//x = append(x, []int{})   if len(nums)==1 改为==0
	fmt.Println(len(x))
	return x
}
func solve(nums []int) [][]int{
	slice:=make([][]int,0)
	//if len(nums)==1 {
	//	temp :=make([]int,len(nums))
	//	copy(temp,nums)
	//	slice = append(slice, temp)
	//	return slice
	//}
	// 修改成这样更简洁
	if len(nums)==0 {
		slice = append(slice, []int{})
		return slice
	}
	x:=solve(nums[1:])
	for i:=0;i< len(x);i++  {
		temp :=make([]int,len(x[i]))
		copy(temp,x[i])
		temp2 :=make([]int,len(x[i]))
		copy(temp2,x[i])
		temp2 = append(temp, nums[0])
		slice = append(slice, temp)
		slice = append(slice, temp2)

	}
	//slice = append(slice,[]int{nums[0]})		// 修改处3 len==0时这句不要了
	return slice

}

func main() {
	fmt.Println(subsets([]int{
		1,2,3,4,
	}))
}
/*
	总结
	1. 这是一道中等递归回溯题
	2. 评论区有大佬用位运算做出来了 0.0 因为对于一个数字有选和不选2种选择 0.0

*/