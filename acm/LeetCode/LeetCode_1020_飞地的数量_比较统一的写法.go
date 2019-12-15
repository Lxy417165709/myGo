package main

var dx [4]int
var dy [4]int

func judge(grid [][]int, x, y int) bool {
	if len(grid) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	if grid[x][y] == 0 {
		return false
	}
	return true
}

func DFS(grid [][]int, x, y int) {
	if judge(grid, x, y) == false {
		return
	}
	grid[x][y] = 0
	for i := 0; i < len(dx); i++ {
		DFS(grid, x+dx[i], y+dy[i])
	}
}
func numEnclaves(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	dx = [4]int{0, 0, -1, 1}
	dy = [4]int{-1, 1, 0, 0}

	m, n := len(A), len(A[0])

	// 去除能到达边界的陆地
	for i := 0; i < m; i++ {
		DFS(A, i, 0)
		DFS(A, i, n-1)
	}
	for i := 0; i < n; i++ {
		DFS(A, 0, i)
		DFS(A, m-1, i)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if judge(A, i, t) {
				ans++
				// 由于统计的是剩余陆地单元格的数量，所以此时我们不需要再DFS消除与其相连的陆地单元格。
			}
		}
	}
	return ans
}

/*
	总结
	1. 这方法是扫描边界，把边界陆地填充为海洋，之后再计算陆地数量。

*/
