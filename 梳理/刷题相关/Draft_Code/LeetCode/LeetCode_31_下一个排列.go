package main

import (
	"fmt"

)

func nextPermutation(nums []int)  {
	flag:=0
	for i:= len(nums)-1;i>=1 && flag==0;i--  {
		if nums[i]>nums[i-1] {
			//num := nums[i]
			index:=i
			for t:=i+1;t<len(nums) ;t++  {
				if nums[t]>nums[i-1] {
					index=t
				}
			}
			nums[i-1],nums[index] = nums[index],nums[i-1]
			//sort.Ints(nums[i:])		// 由于扫描过的区域必然是降序的，所以要升序只需要翻转该切片
			rev(nums[i:])
			flag = 1
		}
	}
	if flag==0 {
		rev(nums[:])
	}
	fmt.Println(nums)
}

func rev(nums []int){
	length:=len(nums)
	for i:=0;i< length/2;i++  {
		nums[i],nums[length-1-i] =nums[length-1-i],nums[i]
	}
}

func main() {
	nextPermutation([]int{2,3,3,1,3})
}


/*
	总结
	1. 该题要一些贪心思维
	2. 一般要求大小的，我们可以考虑下排序
	3. 当数组有序，可通过翻转改变序列方向
	4. 该代码不够优美
*/