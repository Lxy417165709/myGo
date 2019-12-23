package main

func isSymmetric(root *TreeNode) bool {
	queue:=make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue)!= 0{
		if judge(queue)==false{
			return false
		}
		size:=len(queue)

		for i:=0;i<size;i++{
			top:=queue[0]
			queue = queue[1:]
			if top==nil{
				continue
			}
			queue = append(queue,top.Left)
			queue = append(queue,top.Right)
		}
	}
	return true
}

// 判断序列是否回文
func judge(queue []*TreeNode) bool{
	for i:=0;i<len(queue);i++{
		if queue[i]==nil && queue[len(queue)-i-1]==nil{
			continue
		}
		if queue[i]==nil || queue[len(queue)-i-1]==nil{
			return false
		}
		if queue[i].Val!=queue[len(queue)-i-1].Val{
			return false
		}
	}
	return true
}
/*
	总结
	1. 迭代的思路是判断每一层节点是否是回文的，如果每一层都是，则返回true，否则为false
	2. 官方的迭代是和递归思路差不多的，它刚刚开始的时候加入的是2个root节点，详情看官方 0.0
*/