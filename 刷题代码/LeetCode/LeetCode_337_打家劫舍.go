package main

import "fmt"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}
func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}

// BFS DP错误版本
//func rob(root *TreeNode) int {
//
//	queue:=make([]*TreeNode,0)
//	queue = append(queue, root)
//	slice:=make([]int,0)
//	for len(queue)!=0{
//		size:=len(queue)
//		value:=0
//		for i:=0;i<size ;i++  {
//			top:=queue[0]
//			queue = queue[1:]
//			if top==nil {
//				continue
//			}
//			value+=top.Val
//			queue = append(queue, top.Left)
//			queue = append(queue, top.Right)
//
//		}
//		slice = append(slice, value)
//	}
//	fmt.Println(slice)
//	if len(slice)<=2 {
//		if len(slice)==1 {
//			return slice[0]
//		}else{
//			return max(slice[0],slice[1])
//		}
//	}
//
//	dp:=make([]int, len(slice))
//	dp[0] = slice[0]
//	dp[1] = max(dp[0],slice[1])
//	for i:=2;i< len(slice);i++  {
//		dp[i] = max(dp[i-1],dp[i-2]+slice[i])
//	}
//	fmt.Println(dp)
//	return dp[len(slice)-1]
//}
var m map[*TreeNode]int
func rob(root *TreeNode) int {
	m = make(map[*TreeNode]int)
	return solve(root)
	//return max(solve(root,0,0),solve(root,1,0))
}

func solve(root *TreeNode) int{
	if root==nil {
		return 0
	}

	// 减少重复子问题
	if x,ok:=m[root];ok {
		return x
	}


	// 偷
	res1:=root.Val
	if root.Left!=nil {
		res1+=solve(root.Left.Right)+solve(root.Left.Left)
	}
	if root.Right!=nil {
		res1+=solve(root.Right.Right)+solve(root.Right.Left)
	}

	// 不偷
	res2:=solve(root.Left)+solve(root.Right)
	m[root] = max(res1,res2)	// 记得记录
	return m[root]
}

func main() {
	root:=&TreeNode{2,nil,nil}
	root.Left=&TreeNode{1,nil,nil}
	root.Right=&TreeNode{3,nil,nil}
	root.Right.Right=&TreeNode{1,nil,nil}
	root.Left.Left=&TreeNode{4,nil,nil}
	root.Left.Left.Left=&TreeNode{3,nil,nil}

	fmt.Println(rob(root))
}
/*
	总结
	1. 第一次想错了，以为是BFS层序遍历，然后小偷按层偷，但是其实小偷可以偷该层的几户而不会导致警报 0.0.
	2. 看了官方题解后，AC了这题，并且看了他用map去除重叠子问题，自己也写了出来
	3. 总的来说，二叉树的题目要灵活运用 0.0，相关联节点依照题目而定啦 0.0，不过还是逃不过那几层 0.0
*/