package main

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	ans := getHeight(root.Left) + getHeight(root.Right)
	return max(ans, left, right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}
func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	b := max(arr[1:]...)
	a := arr[0]
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 思路是，求取左右子树的直径，之后根节点的直径做比较，取最大的就可以了。
	2. 根节点的直径等于左右子树高度之和
	3. 速度和空间有点慢，应该还有可优化的地方。
*/
