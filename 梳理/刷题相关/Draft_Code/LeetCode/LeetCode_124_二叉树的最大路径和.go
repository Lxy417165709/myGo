package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	mp = make(map[*TreeNode] int)
	ans=-10000000
	search(root)
	return ans
}

var ans int
func search(root *TreeNode){
	if root==nil{
		return
	}
	search(root.Left)
	search(root.Right)
	x:=solve(root,0)
	ans = max(x,ans)
}
var mp map[*TreeNode] int
func solve(root *TreeNode,lay int) int{


	if root==nil{
		return -1000000000
	}

	if lay!=0{
		if x,ok:=mp[root];ok{
			return x
		}
	}

	left:=solve(root.Left,lay+1)
	right:=solve(root.Right,lay+1)
	if lay==0{
		return max(root.Val,root.Val+left,root.Val+right,root.Val+left+right)
	}else{
		mp[root] = max(root.Val,root.Val+left,root.Val+right)
		return mp[root]
	}

}


func max(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}else{
		x := max(arr[:len(arr)-1]...)
		y := arr[len(arr)-1]
		if x>y{
			return x
		}
		return y
	}
}
func main() {

}


/*
	总结
	1. 我的思路是遍历整棵树，让每个树节点都可以作为一个根节点，进行maxPathSum操作
       这样最后就能得到答案了，当然，为了避免重复走，我还用map进行了优化
    2. 看了官方后，官方有更简洁的..但是我还在思考为什么 0.0
	3. 我理解了，他也是把求解和特殊情况分开，return的是求解的,而val里面的就是要求的最大值
	（LeetCode的第二个题解）
*/