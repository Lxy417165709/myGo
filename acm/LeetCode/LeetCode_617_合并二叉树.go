package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return solve(t1,t2)
}

// 返回合并后的二叉树(以第一棵二叉树为基树)
func solve(t1 *TreeNode,t2 *TreeNode) *TreeNode{
	if t1==nil && t2==nil{
		return nil
	}
	if t1==nil{
		return t2
	}
	if t2 == nil{
		return t1
	}

	t1.Val = t1.Val + t2.Val
	left := solve(t1.Left,t2.Left)
	right := solve(t1.Right,t2.Right)
	t1.Left = left
	t1.Right = right
	return t1

}
/*
	总结
	1. 这题目还OK，理清思路就可以AC了，时间复杂度O(n),空间复杂度O(1) (不算上递归栈的空间时)
*/
