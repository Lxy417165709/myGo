package main

func sortedArrayToBST(nums []int) *TreeNode {
	return solve(nums)
}

func solve(nums []int) *TreeNode{
	if len(nums)==0{
		return nil
	}
	mid:=len(nums)>>1
	root := &TreeNode{
		Val: nums[mid],
		Left: solve(nums[:mid]),
		Right:solve(nums[mid+1:]),
	}
	return root
}

/*
	总结
	1. 这题目说左右高度相差不超过1，所以我们只需要把数组分为2半，让他们的个数尽可能相等就可以了
		之后再递归下去。
*/