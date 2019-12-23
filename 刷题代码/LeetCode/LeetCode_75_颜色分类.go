package main

import "fmt"

func sortColors(nums []int)  {
	l:=0
	r:= len(nums)-1
	// [l,r]表示正在处理的空间,curr指向当前需要处理的元素


	for curr:=l;curr<=r;  {
		if nums[curr]==1 {
			curr++
			continue
		}
		if nums[curr]==0 {
			nums[curr],nums[l] = nums[l],nums[curr]
			l++
			curr++	// 第一次错因为这没写 0.0..
			//这curr++是因为我们已经可以保证[0,curr]已经有正确顺序了
			continue
		}
		if nums[curr]==2 {
			nums[curr],nums[r] = nums[r],nums[curr]
			r--
			// 这不curr++，因为数组是 0 2 1 1 2时，curr==1时，交换后,nums[curr] 还是2
			continue
		}
	}
	fmt.Println(nums)
}

func main() {
	sortColors([]int{
		2,0,2,1,1,0,
	})
}

/*
	总结
	1. 这题我想到了双指针法，就是不知道指针移动的方式...
	2. 题解用的是三指针法，l,r,i  思想是只需要把0和2归到一边，1自然就再正确的位置了
	3. 我还没提交 0.0..
	4. 思路我写在上面了
*/