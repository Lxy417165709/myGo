package 杂物

import "fmt"

type TreeNode struct{
	Val int
	Left,Right *TreeNode
}

func main() {
	// 目的: 把root右子树转移到右子树的左子树

	root:= &TreeNode{1,nil,nil}
	root.Right =  &TreeNode{2,nil,nil}
	//root.Right.Left = root.Right	// 这样是不对的,这样操作后，root.Right.Left.Left.Left...无穷无尽
	//fmt.Println(root.Right.Left.Left)
	// &{2 0xc04204e3e0 <nil>}


	root.Right = &TreeNode{3,root.Right,nil}  // 这样就可以了
	fmt.Println(root.Right.Left.Left)
	// <nil>
}

/*

*/
