package main

import "fmt"
//var m map[int]int
//func minDistance(word1 string, word2 string) int {
//	m=make(map[int]int)
//	return solve(word1,word2,0,0)
//}

//// 空间未优化的DP
//func minDistance(word1 string, word2 string) int {
//	//m=make(map[int]int)
//	//return solve(word1,word2,0,0)
//
//	dp :=[1005][1005]int{}
//
//	for i:=1;i<= len(word1);i++  {
//		dp[i][0] = i
//	}
//	for i:=1;i<= len(word2);i++  {
//		dp[0][i] = i
//	}
//	for i:=1;i<=len(word1) ;i++  {
//		for t:=1;t<=len(word2) ;t++  {
//			if word1[i-1]==word2[t-1] {
//				dp[i][t] = dp[i-1][t-1]
//			}else{
//				dp[i][t] = min(dp[i-1][t-1],dp[i-1][t],dp[i][t-1])+1
//			}
//		}
//	}
//	return dp[len(word1)][len(word2)]
//}

// 空间优化的DP...错误的，我不知道怎么搞。。。
func minDistance(word1 string, word2 string) int {
	//m=make(map[int]int)
	//return solve(word1,word2,0,0)

	dp :=[1005]int{}
	for i:=1;i<= len(word2);i++  {
		dp[i] = i
	}
	for i:=1;i<=len(word1) ;i++  {
		last:=i-1
		dp[0]=i
		for t:=1;t<=len(word2) ;t++  {
			tmp:=dp[t]
			if word1[i-1]==word2[t-1] {
				dp[t] = last
			}else{
				dp[t] = min(dp[t-1],dp[t],last)+1
			}
			last=tmp
		}
	}
	return dp[len(word2)]
}


//func solve(word1 string,word2 string,i int,j int) int{
//
//	if x,ok:=m[i<<30|j];ok {
//		return x
//	}
//	if i>= len(word1) && j>=len(word2) {
//		return 0
//	}
//	if i>= len(word1) {
//		return len(word2)-j
//	}
//	if j>= len(word2) {
//		return len(word1)-i
//	}
//
//
//	if word1[i]==word2[j] {
//		m[i<<30|j]=solve(word1,word2,i+1,j+1)
//	}else{
//		m[i<<30|j]= min(
//			solve(word1,word2,i,j+1),
//			solve(word1,word2,i+1,j),
//			solve(word1,word2,i+1,j+1),
//		) + 1
//	}
//	return m[i<<30|j]
//
//
//}
func min(a,b,c int) int{
	if a>b {
		if b>c {
			return c
		}else{
			return b
		}
	}else{
		if a>c {
			return c
		}else{
			return a
		}
	}
}


func main() {
	fmt.Println(minDistance("b",""))
}
/*
	总结
	1. 这是一个DP题，DP动态方程有点难想...我没想出来，我是看官方的 ...
	2. 官方有-1解决，我用+1解决 0.0.表示从尾到iorj的，官方是表示从头开始到iorj的
	3. dp的题，对于数组的设置，可以让数组多一位，让索引0空出来 0.0
	4. 看了另外的题解，然后把空间压缩了。。。在空间这好久，自己做的都错的。。
	5. 滚动数组的话可以参考原来二维数组再设置 0.0
*/