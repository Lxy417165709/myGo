package 二叉树专题

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return solve(root, 1)

}

func solve(root *TreeNode, lay int) int {
	if root == nil {
		return 1000000000
	}

	if root.Left == nil && root.Right == nil {
		return lay
	}
	return int(math.Min(float64(solve(root.Left, lay+1)), float64(solve(root.Right, lay+1))))
}

func main() {
	//root := &TreeNode{1, nil, nil}
	//root.Left = &TreeNode{2, nil, nil}
	//root.Right = &TreeNode{2, nil, nil}
	//root.Right.Left = &TreeNode{3, nil, nil}
	//root.Right.Right = &TreeNode{2, nil, nil}
	//root.Left.Left = &TreeNode{3, nil, nil}
	//root.Left.Right = &TreeNode{3, nil, nil}
	//root.Left.Left.Left = &TreeNode{4, nil, nil}
	//root.Left.Left.Right = &TreeNode{4, nil, nil}
	//root.Left.Left.Left.Left = &TreeNode{4, nil, nil}
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Left = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}


	fmt.Println(minDepth(root))
}

/*
	总结
	1. 这是一道简单的二叉树递归
	2. 注意叶子节点的定义是 没有左右孩子
*/