package main

func isValidBST(root *TreeNode) bool {
	last := -100000000000
	return solve(root,&last)
}

func solve(root *TreeNode,last *int) bool {
	if root == nil{
		return true
	}
	if solve(root.Left,last){
		if root.Val>*last{
			*last = root.Val
			return solve(root.Right,last)
		}
	}
	return false
}

/*
	总结
	1. 这是使用中序遍历的性质，中序遍历的结果是递增的，所以我用last遍历存放最后一个数
		如果满足递增的，则返回true,否则为false
*/