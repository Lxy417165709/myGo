package main


func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{val,nil,nil}
	}

	pre,node := root,root
	for node!=nil{
		pre = node
		if node.Val > val{
			node = node.Left
		}else{
			node = node.Right
		}
	}
	if pre.Val > val{
		pre.Left = &TreeNode{val,nil,nil}
	}else{
		pre.Right = &TreeNode{val,nil,nil}
	}
	return root
}

/*
	总结
	1. 迭代思路和递归差不多，思路是: 创建2个节点，一个指向当前搜寻的节点，一个是其父节点
		通过val和当前节点的值进行的判断，我们可以得出下一个搜索节点，直到搜索节点为nil。
		此时退出循环，操作其父节点，把插入节点插入到该父节点的左子树(小于父节点值时)或右子树(大于父节点值时)
*/



// 去除了pre节点的版本
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{val,nil,nil}
	}

	node := root
	for {
		if node.Val > val{
			if node.Left == nil{
				node.Left = &TreeNode{val,nil,nil}
				break
			}else{
				node = node.Left
			}
		}else{
			if node.Right == nil{
				node.Right = &TreeNode{val,nil,nil}
				break
			}else{
				node = node.Right
			}
		}
	}
	return root
}
