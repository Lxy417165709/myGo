package 二叉树专题

import (
	"fmt"
	"strconv"
)
type TreeNode struct {
	     Val int
	     Left *TreeNode
	    Right *TreeNode
	 }
var ans []string
func binaryTreePaths(root *TreeNode) []string {
	ans = make([]string,0)
	solve(root,"")
	return ans
}

func solve(root *TreeNode,s string){
	if root==nil {
		return
	}
	if root.Right==nil && root.Left==nil {
		s= s + strconv.Itoa(root.Val)
		ans = append(ans, s)
		return
	}

	s= s + strconv.Itoa(root.Val)
	solve(root.Left,s+"->")
	solve(root.Right,s+"->")

	return

}
func main() {
	root := &TreeNode{3, nil, nil}
	//root.Left = &TreeNode{9, nil, nil}
	//root.Left.Left = &TreeNode{10, nil, nil}
	//root.Left.Left.Left= &TreeNode{12, nil, nil}
	//root.Left.Left.Right = &TreeNode{20, nil, nil}
	//root.Right = &TreeNode{20, nil, nil}
	//root.Right.Left = &TreeNode{15, nil, nil}
	//root.Right.Right = &TreeNode{7, nil, nil}


	fmt.Println(binaryTreePaths(root))
}
/*
	总结
	1. 这是一道简单的二叉树题~
*/