package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func pathSum(root *TreeNode, sum int) int {
	if root==nil {
		return 0
	}
	return pathSum(root.Left,sum) + pathSum(root.Right,sum) + solve(root,sum)
}

func solve(root *TreeNode, sum int) int {
	ans:=0
	if root==nil {
		return 0
	}
	if sum-root.Val==0 {
		ans++
	}
	ans+=solve(root.Left,sum-root.Val)
	ans+=solve(root.Right,sum-root.Val)
	return ans
}

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{-3, nil, nil}
	root.Right.Right = &TreeNode{11, nil, nil}

	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Left.Right.Right = &TreeNode{1, nil, nil}

	root.Left.Left.Left  = &TreeNode{3, nil, nil}
	root.Left.Left.Right  = &TreeNode{-2, nil, nil}
	fmt.Println(pathSum(root,8))
}
/*
	总结
	1. 第一次做的时候没思路...虽然是简单题..
	2. 看官方题解发现对每个树求和就可以了..详情看官方题解
*/