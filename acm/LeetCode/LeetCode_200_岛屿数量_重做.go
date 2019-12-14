package main

var dx [4]int
var dy [4]int

func numIslands(grid [][]byte) int {

	dx = [4]int{0, 0, 1, -1}
	dy = [4]int{1, -1, 0, 0}
	ans := 0
	if len(grid) == 0 {
		return ans
	}
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			// 遇到陆地则进行拓展
			if judge(grid, i, t) {
				ans++
				dfs(grid, i, t)
			}
		}
	}
	return ans
}

func judge(grid [][]byte, x, y int) bool {
	if len(grid) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}

	if grid[x][y] == '0' {
		return false
	}
	return true
}

func dfs(grid [][]byte, x, y int) {
	if judge(grid, x, y) == false {
		return
	}
	grid[x][y] = '0'	// 将陆地变为海洋，防止走回头路
	for i := 0; i < len(dx); i++ {
		dfs(grid, x+dx[i], y+dy[i])
	}
}

/*
	总结
	1. 搜索一般模板:
			(1) 确定寻找方向(一般是上下左右)	---->    我们可以采用dx、dy数组
			(2) 防止走回头路					---->	 我们可以修改原矩阵的信息、或者创建一个isVisited映射
			(3) 判断该点是否可以走				---->    可以写一个judge函数，如果该点走过或者越界则返回false
	2. 这题还可以使用广搜，还可以使用并查集
*/
