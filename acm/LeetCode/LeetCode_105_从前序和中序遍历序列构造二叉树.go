package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return solve(preorder,inorder)
}

func solve(preorder []int, inorder []int) *TreeNode{
	if len(preorder)==0{
		return nil
	}

	mid:=preorder[0]
	index:=0
	preLeft:=make([]int,0)
	preRight:=make([]int,0)
	for i:=0;i<len(inorder);i++{
		if inorder[i]==mid{
			index = i
			break
		}
	}

	preLeft = preorder[1:1+index]
	preRight = preorder[1+index:]



	root:=&TreeNode{mid,nil,nil}
	root.Left = solve(preLeft,inorder[:index])
	root.Right = solve(preRight,inorder[index+1:])
	return root


}
func main() {

}

/*
	总结
	1. 第一次AC的时候，采取了很笨的方法获取左右子树的前序遍历，导致时空复杂度过高
	2. 第二次利用性质 （中序遍历的左子树的节点的个数 等于 先序遍历左子树节点的个数，
	   且二者在遍历序列中都是连续的（每一个位置不一定相等，但是所包含的元素是相等的））
*/
