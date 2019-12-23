package main

import (
	"fmt"
)

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

func judge(a int,to_delete []int)bool{
	for i:=0;i< len(to_delete);i++  {
		if a==to_delete[i] {
			return  true
		}

	}
	return false
}
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans = make([]*TreeNode,0)
	//solve(root,to_delete)
	solve(root,nil,-1,to_delete)
	//
	//if judge(root.Val,to_delete) == false{
	//	ans = append(ans, root)
	//}
	return ans
}

var ans []*TreeNode
func show(root *TreeNode){
	if root==nil {
		return

	}

	show(root.Left)
	fmt.Print(root.Val," ")
	show(root.Right)
}
// 这个把所有满足的都加入了，会有重叠 0.0..
func solve(root *TreeNode,father *TreeNode,flag int,to_delete []int){
	if root==nil {
		return
	}else{
		solve(root.Left,root,0,to_delete)
		solve(root.Right,root,1,to_delete)
		if judge(root.Val,to_delete)==false {
			ans = append(ans, root)
		}else{
			if father!=nil {
				if flag==0 {
					father.Left = nil
				}else{
					father.Right = nil
				}
			}
		}
	}
}

//func solve(root *TreeNode, to_delete []int){
//	if root==nil {
//		return
//	}else{
//
//		left,right:=root.Left,root.Right
//		solve(left,to_delete)
//		solve(right,to_delete)
		// false的不管，只管true的，可能是因为false会连通 0.0..
//		if judge(root.Val,to_delete)==true {
//			if left!=nil && judge(left.Val,to_delete)==false {
//				ans = append(ans, left)
//			}
//			if right!=nil && judge(right.Val,to_delete)==false {
//				ans = append(ans, right)
//			}
//		}
//		if left!=nil && judge(left.Val,to_delete)==true {
//			root.Left = nil
//		}
//		if right!=nil && judge(right.Val,to_delete)==true {
//			root.Right = nil
//		}
//
//	}
//
//}


func main() {
	root:=&TreeNode{1,nil,nil}
	root.Left = &TreeNode{2,nil,nil}
	root.Left.Left = &TreeNode{4,nil,nil}
	root.Left.Right = &TreeNode{5,nil,nil}
	root.Right = &TreeNode{3,nil,nil}
	root.Right.Left= &TreeNode{6,nil,nil}
	root.Right.Right = &TreeNode{7,nil,nil}
	delNodes(root,[]int{3,5})
	for i:=0;i< len(ans);i++  {
		show(ans[i])
		fmt.Println()
	}
}
/*

	总结
	1. 第一次做的时候，思路不太清晰 0.0..
	2. 看了些官方题解，可以使用后序遍历 0.0.
*/