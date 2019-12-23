package main

func findTilt(root *TreeNode) int {
	ans = 0
	solve(root)
	return ans
}

var ans int
func solve(root *TreeNode) int{
	if root == nil{
		return 0
	}

	leftSum,rightSum :=solve(root.Left),solve(root.Right)
	ans += abs(leftSum-rightSum)
	return root.Val + leftSum + rightSum
}

func abs(a int) int{
	if a>0{
		return a
	}
	return -a
}
/*
	总结
	1. 这题使用递归很简单，先求出左右子树的和，之后再取其相减绝对值，每一次都这样做就可以了
*/
