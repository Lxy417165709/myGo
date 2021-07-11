package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var arr [][]string
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func printTree(root *TreeNode) [][]string {
	return solve(root)
}
func show(a [][]string) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], len(a[i]))
	}
}

// 获取树高
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}


func solve(root *TreeNode) [][]string {
	arr = make([][]string, 0)
	height := getHeight(root)
	width := 1<<uint(height)-1
	search(root, 0, width/2, 1<<uint(height-2), width, height)
	return arr
}

// 前序遍历填充arr
func search(root *TreeNode, x int, y int, p int, width int, height int) {
	if root == nil {
		return
	}
	if x >= len(arr) {
		arr = append(arr, []string{})
		arr[len(arr)-1] = make([]string, width)
	}
	arr[x][y] = strconv.Itoa(root.Val)
	search(root.Left, x+1, y-p, p/2, width, height)
	search(root.Right, x+1, y+p, p/2, width, height)
}

// 这个方法有bug 0.0... 应该说是错误的 0.0. 按照这个策略得到的矩阵有可能不对称
//func solve(root *TreeNode) [][]string{
//	if root==nil {
//		return [][]string{{""}}
//	}
//	if root.Left==nil && root.Right==nil {
//		return [][]string{
//			{strconv.Itoa(root.Val)},
//		}
//	}
//	left:=solve(root.Left)
//	right:=solve(root.Right)
//	flag:=0
//	if len(left)< len(right) {
//		left,right=right,left
//		flag=1
//	}
//
//
//	tmp:=make([][]string, len(left))
//	for i:=0;i< len(left);i++  {
//		tmp[i] = make([]string,len(left[i]))
//	}
//	for i:=0;i<len(right) ;i++  {
//		base:=(len(tmp[i]) - len(right[i]))/2
//		for t:=0;t<len(right[i]) ;t++  {
//			tmp[i][base + t] = right[i][t]
//		}
//	}
//
//
//	top:=make([]string, len(left[0])*2+1)	// 重点。。。
//	top[len(left[0])] = strconv.Itoa(root.Val)// 重点。。。
//	ans:=make([][]string,0)
//	ans = append(ans,top)
//
//
//	for i:=1;i<=len(left) ;i++  {
//		ans = append(ans,make([]string,0))
//		if flag==0 {
//			ans[i]= append(ans[i], left[i-1]...)
//			ans[i] = append(ans[i],"")
//			ans[i]= append(ans[i], tmp[i-1]...)
//		}else{
//			ans[i]= append(ans[i], tmp[i-1]...)
//			ans[i] = append(ans[i],"")
//			ans[i]= append(ans[i], left[i-1]...)
//		}
//	}
//	return ans
//}


func main() {
	//root := &TreeNode{1, nil, nil}
	//root.Left = &TreeNode{2, nil, nil}
	//root.Right = &TreeNode{3, nil, nil}
	//root.Left.Right = &TreeNode{4, nil, nil}
	show(printTree(nil))
}
/*
	总结
	1. 第一次用递归写出来，但是提交WA了，因为有BUG，得出来的二叉树可能不对称
	2. 过了几天，我用迭代写出来了，思路是先获取树高，通过树高得到树的宽度
       然后通过规律进行迭代填充arr切片(类似矩阵了)
*/