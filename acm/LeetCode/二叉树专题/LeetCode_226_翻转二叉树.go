package 二叉树专题

import "fmt"
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func invertTree(root *TreeNode) *TreeNode {
	return solve(root)
}
func solve(root *TreeNode) *TreeNode{

	if  root==nil{
		return nil
	}
	solve(root.Left)
	solve(root.Right)
	root.Left,root.Right = root.Right,root.Left

	return root
}




func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Left = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}

	x:=invertTree(root)
	fmt.Println(x.Val)
	fmt.Println(x.Right.Right)
	fmt.Println(x.Left.Right.Val)
}
/*
	总结
	1. 这是一道简单的二叉树递归题 0.0..
*/
