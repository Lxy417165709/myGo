package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

var m map[int]int
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {

	m=make(map[int]int)
	solve(root1)
	return solve2(root2,target)

}
func solve(root *TreeNode){
	if root==nil {
		return
	}
	m[root.Val]=1
	solve(root.Left)
	solve(root.Right)

}
func solve2(root *TreeNode,target int) bool{
	if root==nil {
		return false
	}
	if m[target-root.Val]==1 {
		return true
	}

	return solve2(root.Left,target) || solve2(root.Right,target)
}

func main() {
	fmt.Println(twoSumBSTs(nil,nil,5))
}
