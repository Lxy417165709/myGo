package 二叉树专题

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	ans,_ := solve(root)
	return ans
}

// 第一个返回值是这数是否是平衡二叉树，第二个是其树高
// 其实可以用高度-1表示它不是平衡二叉树 0.0
func solve(root *TreeNode) (bool,int) {
	if root==nil{
		return true,0
	}
	leftIsBalanced,leftHeight := solve(root.Left)
	rightIsBalanced,rightHeight := solve(root.Right)

	// 左右子树是平衡二叉树且左右高度差≤1，则该数是平衡二叉树
	return leftIsBalanced && rightIsBalanced && abs(leftHeight-rightHeight)<=1, max(leftHeight,rightHeight)+1
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func abs(a int) int{
	if a>0{
		return a
	}
	return -a
}



/*
	总结
	1. 该题考查平衡二叉树的判断 (通过二叉树左右子数高度)
	2. 官方题解有把获取高度和判断分开为2个函数的方法(暴力法)，我用了一个函数获取高度以及判断
	3. 只求高度就能获取答案了，我们把高度返回值为-1表示不是平衡二叉树就OK了
*/