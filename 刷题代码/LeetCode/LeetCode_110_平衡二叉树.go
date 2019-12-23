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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_,y := solve(root,0)
	return y
}

// 第一个返回值为root层数，第二个返回值为是否为平衡二叉树
func solve(root *TreeNode,lay int) (int,bool){
	if root==nil {
		return lay,true
	}else{
		x1,b1 := solve(root.Left,lay+1)
		x2,b2 := solve(root.Right,lay+1)

		if x1<x2 {
			x2,x1 = x1,x2
		}
		return x1,x1-x2<=1 && b1 && b2
	}
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{2, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	root.Left.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Left.Right = &TreeNode{4, nil, nil}
	root.Left.Left.Left.Left = &TreeNode{4, nil, nil}


	fmt.Println(isBalanced(root))
}

/*
	总结
	1. 该题考查平衡二叉树的判断 (通过二叉树左右子数高度)
	2. 官方题解有把获取高度和判断分开为2个函数的方法(暴力法)，我用了一个函数获取高度以及判断
	3. 只求高度就能获取答案了，我们把高度返回值为-1表示不是平衡二叉树就OK了
*/