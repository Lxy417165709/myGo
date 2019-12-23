package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func hasPathSum(root *TreeNode, sum int) bool {
	return solve(root,sum)
}

func solve(root *TreeNode,sum int) bool{
	if root==nil {
		return false
	}

	if root.Left==nil && root.Right==nil {
		sum = sum - root.Val
		if sum==0 {
			return true
		}else{
			return false
		}
	}

	return solve(root.Left,sum-root.Val) || solve(root.Right,sum-root.Val)
}

func main() {
	root:=&TreeNode{5,nil,nil}
	root.Left = &TreeNode{4,nil,nil}
	root.Right = &TreeNode{8,nil,nil}
	root.Left.Left = &TreeNode{11,nil,nil}
	root.Left.Left.Left = &TreeNode{7,nil,nil}
	root.Left.Left.Right = &TreeNode{2,nil,nil}
	root.Right.Left = &TreeNode{13,nil,nil}
	root.Right.Right = &TreeNode{4,nil,nil}
	root.Right.Right.Right = &TreeNode{1,nil,nil}


	x:=0
	for ;true ;  {
		fmt.Scan(&x)
		fmt.Println(hasPathSum(root,x))
	}
}

/*
	总结
	1.  第一次做的时候是判断是否存在路径(不一定是叶子节点)然后值的和等于target，
		提交失败看案例后，从新审题才发现一定要到叶子，所以重做AC
	2.  该题我用到二叉树递归~
*/