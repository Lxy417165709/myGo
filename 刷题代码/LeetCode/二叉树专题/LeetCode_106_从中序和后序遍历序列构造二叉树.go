package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	return solve(inorder,postorder)
}


func solve(inorder []int, postorder []int) *TreeNode {
	n:=len(postorder)
	if n==0{
		return nil
	}

	mid:=postorder[n-1]
	index:=0

	for i:=0;i<n;i++{
		if inorder[i]==mid{
			index=i
			break
		}
	}


	root := &TreeNode{
		Val: mid,
		Left: solve(inorder[:index], postorder[:index]),
		Right: solve(inorder[index+1:],  postorder[index:n-1]),
	}
	return root
}
func main() {

}

/*
	总结
	1. 和105差不多
*/
