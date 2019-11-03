package main

import "fmt"

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


func flatten(root *TreeNode) {
	solve(root)
}
func solve(root *TreeNode) *TreeNode{
	if root==nil {
		return nil
	}
	if root.Left==nil && root.Right==nil {
		return root
	}

	if root.Left==nil {
		solve(root.Right)
		return root
	}
	if root.Right==nil {
		root.Right = root.Left
		root.Left = nil
		solve(root.Right)
		return root
	}
	left:=solve(root.Left)
	right:=solve(root.Right)

	root.Right = left
	tmp:=root
	for tmp.Right!=nil  {
		tmp = tmp.Right
	}	// 找到最右 0.0.
	tmp.Right = right
	root.Left = nil
	return root
}
func show(root *TreeNode) {
	if root != nil {
		show(root.Left)
		fmt.Println(root.Val)
		show(root.Right)
	}
}
func main() {

	x:=new(TreeNode{1,nil,nil})



	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{5, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}

	root.Left.Left.Left = &TreeNode{10, nil, nil}
	root.Left.Left.Right = &TreeNode{15, nil, nil}
	show(root)
	fmt.Println("----")
	flatten(root)
	show(root)

	fmt.Println("----")
	for root!=nil  {
		fmt.Println(root.Val)
		root = root.Right

	}
}
/*
	总结
	1. 第一次做的时候错了，因为逻辑错了，当时solve函数想写为返回最后的节点，且该树变成了链表
	2. 第二次做对了，修改了之前的判断，逻辑正确就过了 0.0..
	3. 这个模板可以简化思路喔 0.0.
*/