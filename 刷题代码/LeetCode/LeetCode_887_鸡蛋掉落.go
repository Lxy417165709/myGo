package main

var inf int
var isVisit map[int]int

func hash(K int, N int) int {
	return (N << 10) | K
}
func superEggDrop(K int, N int) int {
	inf = 1000000000
	isVisit = make(map[int]int)
	return superEggDropExec(K, N)
}

func superEggDropExec(K int, N int) int {

	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	/*
		if N == 1 {		// 错误1
			return 1
		}
	*/

	hashNumber := hash(K, N)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}
	ans := inf
	for i := 1; i <= N; i++ {
		// 如果掉落碎了
		a := superEggDropExec(K-1, i-1) + 1
		// 如果掉落没碎
		b := superEggDropExec(K, N-i) + 1
		// b := superEggDropExec(K, i - 1) + 1		错误2
		ans = min(ans, max(a, b))
	}
	isVisit[hashNumber] = ans
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 刚刚开始的时候我已经先出了大致框架了，但是状态转移方程和终止条件写错了。上面已经指出了错误的地方。

*/
