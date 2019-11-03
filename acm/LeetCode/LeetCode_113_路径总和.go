package main

import "fmt"

//
type TreeNode struct {
	   Val int
	    Left *TreeNode
	   Right *TreeNode}

var slice [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	slice = make([][]int,0)

	solve(root,sum,[]int{})
	return slice
}


func solve(root *TreeNode,sum int,temp []int){

	if root==nil {
		return
	}
	if root.Left==nil && root.Right==nil {
		if sum-root.Val==0 {
			temp = append(temp, root.Val)
			x := make([]int, len(temp))
			copy(x,temp)
			slice = append(slice, x)

		}
		return
	}
	temp = append(temp, root.Val)
	solve(root.Left,sum-root.Val,temp)
	solve(root.Right,sum-root.Val,temp)
	return
}

func main() {
	root := &TreeNode{5,nil,nil}
	root.Left = &TreeNode{4,nil,nil}
	root.Right = &TreeNode{8,nil,nil}

	root.Left.Left = &TreeNode{11,nil,nil}
	root.Left.Left.Left = &TreeNode{7,nil,nil}
	root.Left.Left.Right = &TreeNode{2,nil,nil}




	root.Right.Left = &TreeNode{13,nil,nil}
	root.Right.Right = &TreeNode{4,nil,nil}
	root.Right.Right.Left= &TreeNode{5,nil,nil}
	root.Right.Right.Right = &TreeNode{1,nil,nil}

	fmt.Println(pathSum(nil,0))

}

/*
	总结
	1. 这是一道简单的二叉树回溯题
	2. 注意自己的例子是否写对0.0..不然代码没错还调半天找哪里出错了...
*/