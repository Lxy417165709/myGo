package main

var mp map[int][]int
var times map[int]int

func treeDiameter(edges [][]int) int {
	mp=make(map[int][]int)
	times=make(map[int]int)
	for i:=0;i<len(edges);i++{
		mp[edges[i][0]] = append(mp[edges[i][0]], edges[i][1])
		mp[edges[i][1]] = append(mp[edges[i][1]], edges[i][0])
		times[edges[i][0]]++
		times[edges[i][1]]++
	}
	ans:=0
	for key,value:=range times{
		if value == 1{
			ans = max(ans,BFS(key))
		}
	}
	return ans
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

//func DFS(a int) int{
//	To := mp[a]
//	for i:=0;i<len(To);i++{
//		DFS(To[i])+1
//	}
//}

// BFS超时了..
func BFS(a int) int{
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
			for t:=0;t< len(To);t++{
				if mark[To[t]]==1{
					continue
				}
				queue = append(queue,To[t])
			}
		}
		level++
	}
	return level
}

func main() {

}
