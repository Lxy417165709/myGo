package main



func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return solve(root,val)
}

func solve(root *TreeNode,val int) *TreeNode{
	if root == nil{
		return &TreeNode{val,nil,nil}
	}
	if val > root.Val{
		root.Right = solve(root.Right,val)
	}else{
		root.Left = solve(root.Left,val)
	}
	return root
}
/*
	总结
	1. 递归写法很简单，比更小就去左子树，否则就去右子树

*/