package main

import "fmt"

const (
	INF = 1000000000000
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var slice []int
func minDiffInBST(root *TreeNode) int {
	slice=make([]int,0)
	mn:=INF
	solve(root)
	for i:=0;i< len(slice)-1;i++  {
		mn = min(mn,slice[i+1]-slice[i])
	}
	fmt.Println(slice)
	return mn
}
// 错误的，最小值并不一定是根与左右 0.0.
//func solve(root *TreeNode) int {
//	if root == nil {
//		return INF
//	}
//	mn := INF
//	if root.Left != nil {
//		mn = min(mn, root.Val-root.Left.Val)
//	}
//	if root.Right != nil {
//		mn = min(mn, root.Right.Val-root.Val)
//	}
//
//	mn = min(mn, solve(root.Left))
//	mn = min(mn, solve(root.Right))
//	return mn
//}

// 下面使用中序遍历解答 0.0.

func solve(root *TreeNode) {
	if root==nil {
		return
	}
	solve(root.Left)
	slice = append(slice, root.Val)
	solve(root.Right)
	return

}

func main() {
	root := &TreeNode{80, nil, nil}
	root.Left = &TreeNode{50, nil, nil}
	root.Left.Left = &TreeNode{10, nil, nil}
	root.Left.Right = &TreeNode{70, nil, nil}
	root.Right = &TreeNode{100, nil, nil}

	fmt.Println(minDiffInBST(root))
}
/*
	总结
	1. 第一次是比较根和左右子节点的差值，发现这样是错误的 0.0 因为还有其他情况，比如根和左子树的右节点0.0
	2. 第二次使用中序遍历，获取切片，空间复杂度O(n)  (最早想到的就是这个方法)
	3. 看了解答后，其实可以把空间优化为O(1),就是在中序遍历的时候，就可以实现判断了 0.0
*/