package main

import (
	"fmt"
)

func permute(nums []int) [][]int {

	ans:=solve(nums)
	fmt.Println(len(ans))
	fmt.Println(test(ans))
	return ans

}

func test(ans [][]int)bool{
	m:=make(map[int]int)
	for i:=0;i<len(ans) ;i++  {
		sum:=0
		for t:=0;t< len(ans[i]);t++  {
			sum = sum*10 + ans[i][t]
		}
		if m[sum]==1 {
			return false
		}else{
			m[sum]=1
		}
	}
	return true

}
func add(origin [][]int,nums[]int) [][]int{
	temp:=make([]int, len(nums))
	copy(temp,nums)
	return append(origin,temp)
}

func solve(nums []int) [][]int{
	ans:=make([][]int,0)
	length := len(nums)
	if length == 1 {

		ans = add(ans, nums)
		return ans
	}

	for i:=length-1;i>=0 ;i--  {

		temp:=make([]int,length)
		copy(temp,nums)
		temp[i],temp[length-1] = temp[length-1],temp[i]
		tempAns:= solve(temp[0:length-1])
		fmt.Println(tempAns)
		for t:=0;t< len(tempAns);t++  {
			tempAns[t] = append(tempAns[t], temp[length-1])
			ans = append(ans, tempAns[t])
		}

	}

	return ans

}

func main() {
	fmt.Println(permute([]int{1,2,3,4,5,6}))
}


/*

	总结
	1. 该题用了回溯法
	2. 尽量不要修改原slice，因为slice是传递指针的，后续对它内容的改变都会影响它之前的内容
	3. length == 2的情况可以省略
*/