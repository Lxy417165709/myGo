package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func countNodes(root *TreeNode) int {
	return solve(root)
}
func solve(root *TreeNode) int{
	if	root==nil{
		return 0
	}
	left,right:=root.Left,root.Right
	l,r:=cnt(left),cnt(right)

	if l==r  {
		return solve(root.Right) + 1<<uint(l)
	}else{
		return solve(root.Left) + 1<<uint(r)
	}
}

func cnt(root *TreeNode) int{
	depth:=0
	for root!=nil{
		root = root.Left
		depth++
	}
	return depth
}




func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Left.Left = &TreeNode{15, nil, nil}
	//root.Left.Right = &TreeNode{7, nil, nil}
	fmt.Println(countNodes(root))
}
/*
	总结
	1. 通过普通的计算节点的方法可以解出这题，但是没有利用到完全二叉树的性质，时间复杂度O(n)
	2. 看了官方的解答后，利用完全二叉树的性质，AC了这题，据说时间复杂度是O(log(n)^2)
	3. 感觉 官方有个用户的 ：利用满二叉树节点公式，分左右子树进行递归。 这个的更规范，他的时间复杂度才是O(log(n)^2) 0.0
	4. 根据他的方法，我重新写了一次 0.0.
*/