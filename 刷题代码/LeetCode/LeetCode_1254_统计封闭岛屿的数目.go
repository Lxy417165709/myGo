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
	if grid[x][y] == 1 {
		return false
	}
	return true
}

func DFS(grid [][]int, x, y int) {
	if judge(grid, x, y) == false {
		return
	}
	grid[x][y] = 1
	for i := 0; i < len(dx); i++ {
		DFS(grid, x+dx[i], y+dy[i])
	}
}

func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dx = [4]int{0, 0, -1, 1}
	dy = [4]int{-1, 1, 0, 0}

	m, n := len(grid), len(grid[0])

	// 去除能到达边界的陆地
	for i := 0; i < m; i++ {
		DFS(grid, i, 0)
		DFS(grid, i, n-1)
	}
	for i := 0; i < n; i++ {
		DFS(grid, 0, i)
		DFS(grid, m-1, i)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if judge(grid, i, t) {
				ans++
				DFS(grid, i, t)
			}
		}
	}
	return ans
}

/*
	总结
	1. 这题本质就是求矩阵中连通块的数目
	2. 注意这题 0表示陆地、1表示海洋。
	3. 这次我是从边界出发DFS的，其实也可以从中间DFS，判断该岛屿是否能到达边界，如果
	   能，那么就不是封闭的，否则就是封闭的。
*/