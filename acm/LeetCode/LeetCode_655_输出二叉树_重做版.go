package main

import "strconv"

var ans [][]string
var width int
var depth int

func printTree(root *TreeNode) [][]string {
	depth = getDepth(root)
	width = (1<<uint8(depth))-1
	p := 1<<uint8(depth-2)  // 注意这里会溢出，但是不会影响最终结果

	ans = make([][]string,0)

	solve(root,0,width/2,p)
	return ans
}
// x是向下的
func solve(root *TreeNode,x int,y int,p int){
	if root == nil{
		return
	}
	if x==len(ans){
		ans = append(ans,make([]string,(1<<uint8(depth))-1))
	}
	ans[x][y] = strconv.Itoa(root.Val)
	solve(root.Left,x+1,y-p,p/2)
	solve(root.Right,x+1,y+p,p/2)

}

func getDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return max(getDepth(root.Left),getDepth(root.Right))+1
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
