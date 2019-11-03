package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	return solve(nums)
}


func solve(nums []int) [][]int{
	ans:=make([][]int,0)
	length := len(nums)
	if length == 1 {
		ans = append(ans, nums)
		return ans
	}

	m := make(map[int]int)	// 与上题一样，由于数组有重复元素，所以要记录该递归层是否使用过某数字
	for i:=length-1;i>=0 ;i--  {
		temp:=make([]int,length)
		copy(temp,nums)
		temp[i],temp[length-1] = temp[length-1],temp[i]
		tailNum:=temp[length-1]
		// 与上一题只有这不同，套了个if
		if m[tailNum]==0 {
			m[tailNum] = 1
			tempAns:= solve(temp[0:length-1])
			for t:=0;t< len(tempAns);t++  {
				tempAns[t] = append(tempAns[t], tailNum)
				ans = append(ans, tempAns[t])
			}
		}
	}

	return ans

}

func main() {
	fmt.Println(permuteUnique([]int{1,1,1,2}))
}


/*

	总结
	1. 该题用了回溯法，在上一题的基础上加了一个map来判断是否使用过某个数字
	2.
*/