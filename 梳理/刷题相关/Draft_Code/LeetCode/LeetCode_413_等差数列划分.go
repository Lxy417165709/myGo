package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	sum:=make([]int,0)
	sum=append(sum, 0)
	s:=0
	for i:=0;i< len(A);i++  {
		s+=A[i]
		sum=append(sum, s)
	}
	ans:=0
	for i:=1;i<=len(A) ;i++  {
		for t:=i+2;t<=len(A) ;t++{
			if 2*(sum[t]-sum[i-1])==(A[i-1]+A[t-1])*(t-i+1) {
				ans++
			}else{
				break
			}
		}
	}
	return ans


}


func main() {
	fmt.Println(numberOfArithmeticSlices([]int{
		1,2,4,6,
	}))
}
/*
	总结
	1. 该题我用了数组前缀和，等差数列公式。。。
	2. 该题目可以使用递归和动态规划
	3. 该题我写的时间复杂度是O(n^2) 0.0.。
*/