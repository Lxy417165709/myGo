package main

import "fmt"

// 法1
//func moveZeroes(nums []int)  {
//	index:=-1
//
//	for i:=0;i< len(nums);i++  {
//		if nums[i]==0 && index ==-1 {
//			index = i
//		}else{
//			if nums[i]!=0 && index!=-1 {
//				nums[index],nums[i] = nums[i],nums[index]
//				i = index
//				index = -1
//			}
//		}
//	}
//	//fmt.Println(nums)
//}

// 法2
func moveZeroes(nums []int)  {
	index:=0

	for i:=0;i< len(nums);i++  {
		if nums[i]!=0 {
			nums[index] = nums[i]

			// 这一步没有就会出错!
			if i!=index {
				nums[i] = 0
			}
			index++
		}
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{1,3})
}

/*
	总结
	1. 法1比法2复杂，时间花费更大，因为法1用的是交换思维，而法2用的是替代思维
	2. 注意细节处理
	3. 官方题解的最优解是法1的优化
*/