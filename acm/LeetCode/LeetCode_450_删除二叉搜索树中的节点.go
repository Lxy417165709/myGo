package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	return solve(root, key)
}

func solve(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}


	// 优化步骤，根据二叉搜索树的性质
	if key<root.Val{
		root.Left = solve(root.Left, key)
	}else{
		root.Right = solve(root.Right, key)
	}
	//root.Left = solve(root.Left, key)
	//root.Right = solve(root.Right, key)
	if root.Val == key {

		if root.Left == nil {
			root = root.Right
		} else {
			right := root.Right
			root = root.Left
			tmp := root
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = right
		}
	}
	return root
}
func show(root *TreeNode) {
	if root != nil {
		show(root.Left)
		fmt.Println(root.Val)
		show(root.Right)
	}
}

func main() {
	//root := &TreeNode{5, nil, nil}
	//root.Left = &TreeNode{3, nil, nil}
	//root.Right = &TreeNode{6, nil, nil}
	//root.Left.Left = &TreeNode{2, nil, nil}
	//root.Left.Left.Left = &TreeNode{2, nil, nil}
	//root.Left.Left.Left.Left = &TreeNode{2, nil, nil}
	//root.Left.Right = &TreeNode{4, nil, nil}
	//root.Right.Right = &TreeNode{7, nil, nil}
	tmp := deleteNode(nil, 5)
	show(tmp)
}
/*
	总结
	1. 第一次没优化，直接全树找，之后才想到二叉搜索树的性质，优化代码在上面了
	2. 找到节点后的替代操作有多种，我采用的是右子树插入左子树的最右 然后左子树代替根 的方法
	3. 官方还有其他的替代方法，比如找到左子树的最大值代替，右子树的最小值代替
	4. 这道题没有想象的难 0.0..
*/