package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if rootIsSelfOrNil(root, p, q) {
		return root
	}
	lcaSearchedInLeftTree := lowestCommonAncestor(root.Left, p, q)
	lcaSearchedInRightTree := lowestCommonAncestor(root.Right, p, q)
	switch {
	case lcaIsRoot(lcaSearchedInLeftTree, lcaSearchedInRightTree):
		return root
	case lcaIsExistInLeftTree(lcaSearchedInLeftTree, lcaSearchedInRightTree):
		return lcaSearchedInLeftTree
	case lcaIsExistInRightTree(lcaSearchedInLeftTree, lcaSearchedInRightTree):
		return lcaSearchedInRightTree
	}
	return nil
}

func rootIsSelfOrNil(root, p, q *TreeNode) bool {
	return root == nil || root == p || root == q
}

func lcaIsRoot(lcaSearchedInLeftTree, lcaSearchedInRightTree *TreeNode) bool {
	return lcaSearchedInLeftTree != nil && lcaSearchedInRightTree != nil
}

func lcaIsExistInLeftTree(lcaSearchedInLeftTree, lcaSearchedInRightTree *TreeNode) bool {
	return lcaSearchedInLeftTree != nil && lcaSearchedInRightTree == nil
}
func lcaIsExistInRightTree(lcaSearchedInLeftTree, lcaSearchedInRightTree *TreeNode) bool {
	return lcaSearchedInLeftTree == nil && lcaSearchedInRightTree != nil
}
