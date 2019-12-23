package main

var lastVal int

func convertBST(root *TreeNode) *TreeNode {
	lastVal = 0
	convertBSTExec(root)
	return root
}

func convertBSTExec(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTExec(root.Right)
	root.Val += lastVal
	lastVal = root.Val
	convertBSTExec(root.Left)
}

/*
	总结
	1. 这题其实就是反序的中序遍历。
	2. 我的思路是，通过一个外部变量，记录大于当前节点的总和，那么当前节点的值就会等于 原值+该总和
	   之后总和再变为当前节点的值，继续进行递归。
*/
