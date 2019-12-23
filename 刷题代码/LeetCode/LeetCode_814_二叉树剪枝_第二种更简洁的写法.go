package main

func pruneTree(root *TreeNode) *TreeNode {

	return solve(root)
}



// 第二种解法
func solve(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	root.Left = solve(root.Left)
	root.Right = solve(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0{
		return nil
	}
	return root
}
/*
	总结
	1. 这个代码更简洁了 0.0
*/
