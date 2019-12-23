package main
var mp map[int][]int

func treeDiameter(edges [][]int) int {
	if len(edges)==0{
		return 0
	}
	mp=make(map[int][]int)
	root:=0
	for i:=0;i<len(edges);i++{
		root =edges[i][0]
		mp[edges[i][0]] = append(mp[edges[i][0]], edges[i][1])
		mp[edges[i][1]] = append(mp[edges[i][1]], edges[i][0])
	}
	ans:=0
	ans = max(ans,BFS(root,&root))
	ans = max(ans,BFS(root,&root))
	return ans
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func BFS(a int,depthestNode *int) int{
	mark:=make(map[int]int)
	queue:=make([]int,0)
	queue = append(queue, a)
	mark[a] = 1
	level:=0
	for len(queue)!=0{
		size:=len(queue)

		for i:=0;i<size;i++{
			top:=queue[0]
			queue = queue[1:]
			To:=mp[top]

			*depthestNode = top
			for t:=0;t< len(To);t++{
				if mark[To[t]]==1{
					continue
				}
				mark[To[t]]=1
				queue = append(queue,To[t])
			}
		}
		level++
	}
	return level-1
}

/*
	总结
	1. 第一次做的时候采用了半暴力的方式，每次从叶子节点开始找最大的路径，这方法超时了
	2. 看了官方题解后才明白，我们只需要先找出一个最深的点，之后从这个点出去再找最长的路径，那么这条路径就是我们要求的 0.0

*/
