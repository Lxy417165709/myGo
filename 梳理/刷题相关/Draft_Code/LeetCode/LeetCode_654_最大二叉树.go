package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return solve(nums)
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func solve(nums []int) *TreeNode{
	if len(nums)==0{
		return nil
	}
	mx:=nums[0]
	index:=0
	for i:=0;i<len(nums);i++{
		if nums[i]>mx{
			mx = nums[i]
			index = i
		}
	}

	root := &TreeNode{
		Val: mx,
		Left: solve(nums[:index]),
		Right: solve(nums[index+1:]),
	}
	return root

}

// 迭代实现
func solve(nums []int) *TreeNode{
	if len(nums)==0{
		return nil
	}
	root:=&TreeNode{nums[0],nil,nil}
	for i:=1;i<len(nums);i++{

		if nums[i]>root.Val{
			root = &TreeNode{nums[i],root,nil}
		}else{
			tmp:=root
			for tmp.Right!=nil && tmp.Right.Val>nums[i]{
				tmp = tmp.Right
			}
			tmp.Right = &TreeNode{nums[i],tmp.Right,nil}
		}
	}
	return root
}


func main() {

}

/*
	总结
	1. 该题我是使用递归实现的，但是官方还有用迭代实现的算法
	2. 思路是从左到右遍历数组，根据一定的规则调整树的结构
*/
