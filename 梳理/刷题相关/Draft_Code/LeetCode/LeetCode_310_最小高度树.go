package main

import "fmt"

//var information [][]int
//var m map[int]int
//func min(a, b int) int{
//	if	a<b{
//		return a
//	}
//	return b
//}
//func BFS(x int) int{
//	m = make(map[int]int)
//	queue :=make([]int,0)
//	queue = append(queue, x)
//	m[x]=1
//	height:=0
//	for ; len(queue)!=0;  {
//		size:= len(queue)
//		height++
//		for k:=1;k<=size ;k++  {
//			top:=queue[0]
//			queue=queue[1:]
//			for i:=0;i< len(information[top]);i++  {
//				point:=information[top][i]
//				if m[point]==0 {
//					queue = append(queue, point)
//					m[point]=1
//				}
//			}
//		}
//	}
//	return height
//
//}
//
//func findMinHeightTrees(n int, edges [][]int) []int {
//	if n==0 {
//		return []int{}
//	}
//	information=make([][]int,n)
//	for i:=0;i< len(edges);i++  {
//		x,y:=edges[i][0],edges[i][1]
//		information[x] = append(information[x],y)
//		information[y] = append(information[y],x)
//	}
//	slice:=make([]int,0)
//	mn:=1000000000
//	for i:=0;i< n;i++  {
//		ans:=BFS(i)
//		//fmt.Println(i,ans)
//		slice = append(slice, ans)
//		mn=min(mn,ans)
//	}
//	ans:=make([]int,0)
//	for i:=0;i<n ;i++  {
//		if slice[i]==mn {
//			ans = append(ans, i)
//		}
//	}
//	return ans
//}
func findMinHeightTrees(n int, edges [][]int) []int {
	if n==1 {
		return []int{0}
	}


	information:=make([][]int,n)// information[i]表示与i相连的点的编号
	degree :=make([]int,n)		//degree[i]表示 点i的度
	for i:=0;i< len(edges);i++  {
		x,y:=edges[i][0],edges[i][1]
		information[x] = append(information[x],y)
		information[y] = append(information[y],x)
		degree[x]++
		degree[y]++
	}
	queue:=make([]int,0)
	for i:=0;i<n ;i++  {
		if degree[i]==1 {
			queue = append(queue, i)
		}
	}
	all :=n


	for ; all>2;  {
		// 和第一次AC的只有这片中括号区域不同
		all -=len(queue)
		length:=len(queue)
		for i:=0;i< length;i++  {
			points:=information[queue[0]]
			queue = queue[1:]	// 出队
			for t:=0;t< len(points);t++  {
				degree[points[t]]--
				x:=degree[points[t]]
				if x==1 {
					queue = append(queue, points[t])
				}
			}
		}
	}
	return queue
}
func main() {
	fmt.Println(findMinHeightTrees(0,[][]int{

	}))
}

/*
	总结
	1. 第一次用BFS做，超时了 0.0.
	2. BFS注意点，起始点记得要入队标记 0.0..
	3. 看了官方题解，他们采用了分层剥削，每次都把叶子节点去除 0.0 晚点再AC这题 0.0
	4. 解答了这题了，其实还是在纠结用什么数据结构 0.0..然后想到了感觉时间也很多，于是没实现，于是借鉴官方AC这题
		(官方采用的数据结构我也有想过，只是我不知道怎么优美的运用 0.0.)
	5. 该题用队列结构+数组最好 0.0
	6. 第一次AC提交的版本没用到队列 想看可以回去力扣看
	7. 分层的可以考虑BFS喔 0.0..
*/