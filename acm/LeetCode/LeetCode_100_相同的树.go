package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil && q==nil {
		return true
	}
	if p==nil||q==nil {
		return false
	}

	return p.Val==q.Val && isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)


	
}

func main() {

}


/*

	总结
	1. 该题考查二叉树递归
	2. 递归不一定是一棵树的事，还可以使用多颗喔~
*/