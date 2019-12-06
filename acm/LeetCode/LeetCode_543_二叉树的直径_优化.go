package main



var ans int
func diameterOfBinaryTree(root *TreeNode) int {
	ans=0
	getHeight(root)
	return ans
}

func getHeight(root *TreeNode) int{
	if root == nil{
		return 0
	}
	left,right:=getHeight(root.Left),getHeight(root.Right)
	ans = max(ans,left+right)
	return max(left,right)+1
}
func max(arr ...int) int{
	if len(arr) == 1{
		return arr[0]
	}
	b:=max(arr[1:]...)
	a:=arr[0]
	if a>b{
		return a
	}
	return b
}
/*
	总结
	1. 其实在求树高的时候，我们就已经可以得出直径了，所以此时我们用一个全局变量ans
		记录下求树高过程中的最大直径，之后输出就可以了。
*/