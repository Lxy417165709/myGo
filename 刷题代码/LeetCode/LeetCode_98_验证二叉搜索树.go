package main

import (
	"fmt"
)
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func isValidBST(root *TreeNode) bool {
	min,max,flag:=solve(root)
	fmt.Println(min,max,flag)
	return flag

}
func mn(a,b int) int{
	if a>b {
		return b
	}
	return a
}

func mx(a,b int) int{
	if a>b {
		return a
	}
	return b
}

func solve(root *TreeNode) (min int,max int,flag bool){
	if root==nil {
		return -1000000000,1000000000,true
	}

	if root.Left==nil && root.Right==nil {
		return root.Val,root.Val,true
	}
	if root.Left==nil{
		a,b,c:=solve(root.Right)
		if !(a>root.Val) {
			c = false
		}
		a = mn(a,root.Val)
		b = mx(b,root.Val)
		return a,b,c
	}
	if root.Right==nil{
		a,b,c:=solve(root.Left)
		if !(b<root.Val) {
			c=false
		}
		a = mn(a,root.Val)
		b = mx(b,root.Val)
		return a,b,c
	}
	a,b,c:=solve(root.Left)
	d,e,f:=solve(root.Right)
	if root.Val>b && root.Val<d {
		return mn(a,root.Val),mx(e,root.Val),c&&f
	}else{
		return mn(a,root.Val),mx(e,root.Val),false
	}
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{1, nil, nil}
	//root.Right = &TreeNode{10, nil, nil}
	//root.Right.Left = &TreeNode{11, nil, nil}
	//root.Right.Right = &TreeNode{11, nil, nil}
	isValidBST(root)
}

/*
	总结
	1. 第一次错的时候，是由于没看清楚题，题目要求的这二叉搜索树，左一定小于根，
       右一定大于根，和普通的二叉搜索树不太一样，差别在于普通的是左一定小于等于根
	2. 第一次想到的是中序遍历，感觉很简单，为了锻炼递归，就写了递归了 0.0.
*/