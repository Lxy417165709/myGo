package 二叉树专题

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	//if root==nil {
	//	return true
	//}
	//return solve(root.Left,root.Right)

	// 这样写更简洁
	return solve(root, root)
}

// 该函数判断两棵树是否镜像对称
func solve(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && solve(left.Right, right.Left) && solve(left.Left, right.Right)

}

func show(root *TreeNode) {
	if root != nil {
		show(root.Left)
		fmt.Println(root.Val)
		show(root.Right)
	}
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{2, nil, nil}

	c := root
	a := root.Left
	b := root.Right
	root = a
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root = b
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{4, nil, nil}

	root = c
	show(root)
	fmt.Println(isSymmetric(root))
}

/*
	总结
	1. 该题考查二叉树的对称判断，需要用到递归
	2. 官方题解还有使用迭代的.

*/
