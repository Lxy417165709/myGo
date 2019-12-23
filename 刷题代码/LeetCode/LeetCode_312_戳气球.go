package main

var off uint8
var coins map[int]int

func maxCoins(nums []int) int {
	coins = make(map[int]int)
	off = 15
	return maxCoinsExec(nums, 0, len(nums)-1)
}

func maxCoinsExec(nums []int, l, r int) int {
	hashNumber := r<<off | l
	if x, ok := coins[hashNumber]; ok {
		return x
	}
	ans := 0
	for i := l; i <= r; i++ {
		// 这里纠结了好久，因为无法确定到底i的左边是哪个数(因为有可能i左边这个数已经被取了)
		getCoins := nums[i]

		// 在刚刚开始可以构造nums为头尾为1，之后l,r从1和len(nums)-1开始遍历，那就不需要这里的判断了
		// 如果刚刚开头没有，那么就要判断是否越界了
		// 这里可能会问，为什么i的旁边不是i-1和i+1呢？
		// 		这是因为i-1可能已经是被选取了的，那么去除了选取后的，i的左边就是l-1,同理可以得到右边为r+1
		if l-1 >= 0 {
			getCoins *= nums[l-1]
		}
		if r+1 <= len(nums)-1 {
			getCoins *= nums[r+1]
		}
		ans = max(ans, maxCoinsExec(nums, l, i-1)+maxCoinsExec(nums, i+1, r)+getCoins)
	}
	coins[hashNumber] = ans
	return coins[hashNumber]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 这题好像属于区间dp，数学表示就是，在一个数组中选取一个数后，将该数和其左右两表的乘积加入结果(如果旁边没数，那么数==1)
	   求最大的结果。
	2. 刚刚开始我是把nums划分为子nums再合并的，但是这样就用不了记忆化搜索了，之后看了官方题解后，才有了上面的解决方案。
	3. 这题还没理解清楚
*/
