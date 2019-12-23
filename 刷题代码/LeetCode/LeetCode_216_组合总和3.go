package main

func combinationSum3(k int, n int) [][]int {
	return solve(k, n, 1)
}

func newSlice(slice []int) []int {
	tmp := make([]int, len(slice))
	copy(tmp, slice)
	return tmp
}

func solve(k int, n int, low int) [][]int {
	if k == 1 {
		if n >= low && n <= 9 {
			return [][]int{{n}}
		} else {
			return [][]int{}
		}
	}
	ans := [][]int{}
	for i := low; i <= 9; i++ {
		after := solve(k-1, n-i, i+1)
		for t := 0; t < len(after); t++ {
			tmp := newSlice(after[t])
			tmp = append(tmp, i)
			ans = append(ans, tmp)
		}
	}

	return ans
}

/*
	总结
	1. 该题我使用了回溯法，根据题目的限制条件进行回溯，提交后AC了
	2. 官方预期结果 k=1 n>=10时，此时给的结果是错的
*/