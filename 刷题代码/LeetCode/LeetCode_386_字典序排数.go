package main

import "fmt"

var ans []int
var m map[int]int
var i int
// 正确策略
//func solve(num int,n int) {
//	i++
//	// 这个优化和没优化一样 0.0..因为没有重叠子问题 0.0..
//	if _,ok:=m[num];ok {
//		return
//	}
//	m[num]=1
//
//	if num>n {
//		return
//	}
//	ans = append(ans, num)
//	num = num*10
//	for i:=0;i<10 ;i++  {
//		// 这个剪枝优化使得i变小了，但是时间还是那样...
//		if num+i>n{
//			return
//		}
//
//		solve(num+i,n)
//	}
//
//}

func solve(n int){
	now:=1
	for i:=0;i<n ;i++  {

		ans = append(ans, now)

		if now*10<=n {
			now = now*10
		}else{
			// 这个>= 等于是因为在上一次循环中，我们已经加入now了，如果没有等于 下一次加入的是now+1可能会超过n 0.0
			if now>=n {
				now = now/10
			}
			now++
			for now%10==0{
				now = now/10
			}
		}
	}

}

func lexicalOrder(n int) []int {
	//i=0
	//m=make(map[int]int)
	ans=make([]int,0)
	//for i:=1;i<10 ;i++  {
	//	solve(i,n)
	//}
	//fmt.Println(i)

	solve(n)	// 迭代写法
	fmt.Println(len(ans))
	return ans
}
// 错误策略
//var ans []int
//var m map[int]int
//func solve(now int,n int){
//	if _,ok:=m[now];ok {
//		return
//	}
//	m[now] = 1
//	if now<=n {
//		ans = append(ans, now)
//	}else{
//		return
//	}
//	solve(now*10,n)
//	for i:=1;i<=10 ;i++  {
//
//		if now%10==0 {
//			solve((now+i)/10,n)
//		}else{
//			solve(now+i,n)
//		}
//	}
//
//}



func main() {
	fmt.Println(lexicalOrder(213))
}
/*
	总结
	1. 第一次采用了错误的递归方式，然后WA了
	2. 之后改用迭代法，但是还是WA，因为策略也错了，此时摸出了一些规律，想到了递归
	3. 偷了下懒，看了官方的..发现有用DFS和迭代的，然后自己就用DFS写出来了0.0.
	4. 官方有用迭代前序遍历的，也有就使用迭代的...我膜拜了...
*/