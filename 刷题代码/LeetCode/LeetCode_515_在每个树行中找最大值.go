package main

import "fmt"

type TreeNode struct {
	     Val int
	Left *TreeNode
	     Right *TreeNode
	 }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var ans []int
func largestValues(root *TreeNode) []int {
	ans = make([]int,0)
	solve(root,0)
	return ans
}

 func max(a,b int) int{
	 if a>b {
		 return a

	 }
	 return b
 }
func solve(root *TreeNode,lay int) {
	if root==nil {
		return
	}
	if len(ans)<lay+1 {
		ans = append(ans, root.Val)
	}else{
		ans[lay] = max(ans[lay],root.Val)
	}
	solve(root.Left,lay+1)
	solve(root.Right,lay+1)


}


func main() {
	root := &TreeNode{1, nil, nil}
	//root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	//root.Left.Left = &TreeNode{5, nil, nil}
	//root.Left.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}
	fmt.Println(largestValues(root))
}
/*
	总结
	1. 该题我使用前序遍历法，记录每个节点的层，然后和该层最大值做比较，之后写回数组，然后就OK了
	2. 其实可以用迭代的方式实现，可以先获取每个层的所有节点，再一一比较取得最大值0.0.
*/