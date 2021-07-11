package main

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil{
		return []int{}
	}
	vector := [505][]int{}
	queue:=[]*TreeNode{}
	queue = append(queue,root)
	for len(queue)!=0{
		top:=queue[0]
		queue = queue[1:]

		if top.Left!=nil{
			vector[top.Val] = append(vector[top.Val],top.Left.Val)
			vector[top.Left.Val] = append(vector[top.Left.Val],top.Val)
			queue = append(queue,top.Left)
		}
		if top.Right!=nil{
			vector[top.Val] = append(vector[top.Val],top.Right.Val)
			vector[top.Right.Val] = append(vector[top.Right.Val],top.Val)
			queue = append(queue,top.Right)
		}
	}

	mark := [505]bool{}
	myQueue := []int{}
	myQueue = append(myQueue,target.Val)
	mark[target.Val]=true
	ans := []int{}

	nowTimes:=0
	for len(myQueue)!=0{
		size:=len(myQueue)
		for i:=0;i<size;i++{
			top:=myQueue[0]
			myQueue = myQueue[1:]

			if nowTimes==K{
				ans = append(ans,top)
				continue
			}
			for t:=0;t<len(vector[top]);t++{
				if mark[vector[top][t]]==true{
					continue
				}
				myQueue = append(myQueue,vector[top][t])
				mark[vector[top][t]]=true
			}
		}
		nowTimes++
	}


	return ans
}
/*
	总结
	1. 这题目我是把这棵树变成图，然后遍历距离为K的节点，这样就可以AC了
	2. 但我觉得有更优的算法
*/
