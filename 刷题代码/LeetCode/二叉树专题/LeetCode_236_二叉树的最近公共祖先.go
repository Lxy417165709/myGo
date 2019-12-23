package 二叉树专题

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	_,Sp :=search(root,p,"")
	_,Sq :=search(root,q,"")
	prefix:=getPrefix(Sp,Sq)
	fmt.Println(Sp,Sq,prefix)

	return prefixToTreeNode(root,prefix)
}
func prefixToTreeNode(root *TreeNode,prefix string) * TreeNode{
	if prefix=="" {
		return root
	}else{
		if prefix[0]=='0' {
			return prefixToTreeNode(root.Left,prefix[1:])
		}else{
			return prefixToTreeNode(root.Right,prefix[1:])
		}
	}
}
// 第一次错在这，公共前缀的获取，当不相等时应该直接退出
func getPrefix(a,b string)string{
	ans:=""
	for i:=0;i< len(a) && i< len(b) && a[i]==b[i];i++  {
		ans+=string(a[i])
	}
	return ans
}

func search(root *TreeNode,target *TreeNode,markStr string) (bool,string){
	if root==nil {
		return false,markStr
	}
	if root.Val == target.Val{
		return true,markStr
	}else{
		Lflag,left:=search(root.Left,target,markStr+"0")
		Rflage,right:=search(root.Right,target,markStr+"1")

		if Lflag==true {
			return Lflag,left
		}else{
			return Rflage,right
		}

	}
}



func main() {
	root :=&TreeNode{3,nil,nil}
	root.Left = &TreeNode{5,nil,nil}
	root.Left.Left = &TreeNode{6,nil,nil}
	root.Left.Right = &TreeNode{2,nil,nil}
	root.Left.Right.Left = &TreeNode{7,nil,nil}
	root.Left.Right.Right = &TreeNode{4,nil,nil}
	root.Right = &TreeNode{1,nil,nil}
	root.Right.Left = &TreeNode{0,nil,nil}
	root.Right.Right = &TreeNode{8,nil,nil}

	fmt.Println(lowestCommonAncestor(root,root.Left.Left,root.Right.Left))
}
/*
	本题目已经说明:
    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉树中。

	总结
	1. 思路：标记到pq的路径字符串，通过2个字符串获取公共前缀，再通过公共前缀到二叉树中找这个节点
	2. 第一次错在公共前缀的获取，公共前缀的获取，当不相等时应该直接退出，不用判断后面的字符了
	3. 官方有用递归迭代解决的 0.0.我这个和他们方法不一样，虽然直观，但是代码多 0.0..
*/