package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	return solve(nums)
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

		ans = append(ans,nums)	// 这句会把nums的引用传递给ans[0],当nums改变时，ans[0]也会改变..
								// 因为nums没变过，所以可以直接append
		//ans = add(ans, nums)	// 这是我自己写的函数，这样就不会引用传递，代码如上
		return ans
	}

	for i:=length-1;i>=0 ;i--  {
		temp:=make([]int,length)
		copy(temp,nums)			// 为了防止原切片被修改，这使用了一个temp切片存放原切片的元素
		temp[i],temp[length-1] = temp[length-1],temp[i]	// 交换，固定最后一位
		tempAns:= solve(temp[0:length-1])				// 获得[0,length-2]索引的全排列
		tailNum:=temp[length-1]							// 这是固定的一位(我固定的是最后一位)

		// 合并
		for t:=0;t< len(tempAns);t++  {
			tempAns[t] = append(tempAns[t], tailNum)
			ans = append(ans, tempAns[t])
		}

	}

	return ans

}

func main() {
	fmt.Println(permute([]int{1,2,3}))
}


/*

	总结
	1. 该题用了回溯法
	2. 尽量不要修改原slice，因为slice是传递指针的，后续对它内容的改变都会影响它之前的内容
	3. length == 2的情况可以省略
*/