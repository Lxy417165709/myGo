package 二叉树专题

import "fmt"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func sumOfLeftLeaves(root *TreeNode) int {
	if root==nil {
		return 0
	}
	return solve(root.Left,1) + solve(root.Right,2)
}

func solve(root *TreeNode,mark int) int{
	if root==nil {
		return 0
	}

	if root.Left==nil && root.Right==nil {
		if mark==1 {
			return root.Val
		}
	}
	return solve(root.Left,1) + solve(root.Right,2)
}
func main() {
	//root:=&TreeNode{3,nil,nil}

	fmt.Println(sumOfLeftLeaves(nil))
}
/*

	总结
	1. 这是一道简单的二叉树递归题~
*/