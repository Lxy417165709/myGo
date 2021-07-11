package 二叉树专题

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


 var i int
 var ans int
 // 题目说了此树非空
func kthSmallest(root *TreeNode, k int) int {
	// 注意初始化
	i=1
	ans=0
	solve(root,k)
	return ans
}

func solve(root *TreeNode,k int) {
	if root==nil {
		return
	}
	solve(root.Left,k)
	// i表示现在找到了第i小的
	if k==i {
		ans = root.Val
	}
	i++
	solve(root.Right,k)
	return
}
func main() {
	root:=&TreeNode{100,nil,nil}
	root.Left = &TreeNode{3,nil,nil}
	root.Left.Right = &TreeNode{4,nil,nil}
	root.Left.Left = &TreeNode{2,nil,nil}
	root.Left.Left.Left = &TreeNode{1,nil,nil}
	root.Right = &TreeNode{610,nil,nil}
	fmt.Println(kthSmallest(root,6))

}
/*
	总结
	1. 第一次想的时候，想到了中序遍历，思路是先中序遍历出所有元素的序列(已排序)，之后返回第index = k-1的元素就OK了
	2. 看了官方题解，他也是用中序遍历，但是在找到第k小的元素后就返回，方法可以看上面
*/