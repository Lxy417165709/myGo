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

func maxDepth(root *TreeNode) int {
	return solve(root,0)
}

func solve(root *TreeNode, lay int) int {
	if root==nil {
		return lay
	}else{
		return int(math.Max(float64(solve(root.Left,lay+1)),float64(solve(root.Right,lay+1))))
	}
}

func main() {

}


/*
	总结
	1. 该题考查如何获取二叉树深度
*/