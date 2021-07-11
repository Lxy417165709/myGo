package main

import "fmt"

func productExceptSelf(nums []int) []int {
	sumL,sumR:=1,1
	ans:=make([]int,0)
	for i:=0;i< len(nums);i++  {
		ans = append(ans, 1)
	}
	for i:=0;i< len(nums);i++  {
		ans[i] = ans[i] * sumL
		sumL = sumL * nums[i]
	}
	for i:=len(nums)-1;i>= 0;i--  {
		ans[i] = ans[i] * sumR
		sumR = sumR * nums[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{
		1,2,
	}))
}


/*
	总结
	1. 有时候2个问题可以分别解决了1个后再解决另外一个
       在该题中 x[i] = A[0]..A[i-1] * A[i+1]...*A[n-1]
	   那我们可以先把从左到右 A[0]..A[i-1]算入x[i]
       之后我们再从右到左 把 A[n-1]..A[i+1]算入x[i]
*/