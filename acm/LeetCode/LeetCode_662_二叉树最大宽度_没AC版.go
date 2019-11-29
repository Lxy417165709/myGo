package main

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

var treeArr [][]bool
func widthOfBinaryTree(root *TreeNode) int {
	treeArr = make([][]bool,0)
	solve(root,0,0)
	mx:=0
	for i:=0;i<len(treeArr);i++{
		// fmt.Println(treeArr[i])
		l,r:=0,(1<<uint8(i))-1
		for l<=r && treeArr[i][l]==false{
			l++
		}
		for l<=r && treeArr[i][r]==false{
			r--
		}
		mx = max(mx,r-l+1)
	}
	return mx
}

func solve(root *TreeNode,lay int,number int){
	if root == nil{
		return
	}
	if len(treeArr)==lay{
		treeArr = append(treeArr,make([]bool,1<<uint8(lay)))
	}
	treeArr[lay][number] = true
	solve(root.Left,lay+1,number*2)
	solve(root.Right,lay+1,number*2+1)
}
/*
	总结
	1. 这个版本超出内存了...
	2. 思路是通过树，构建出其数组形式，然后得到每一层最左和最右的树节点，之后就可以得出宽度了，然后找出最大的宽度，
*/
