package main

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if grid[i][t] == 1 {
				for u := 0; u < len(dx); u++ {
					ans += getWeight(grid, i+dx[u], t+dy[u])
				}
			}
		}
	}
	return ans
}

func getWeight(grid [][]int, x, y int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return 1
	}
	return grid[x][y] ^ 1
}

/*
	总结
	1. 这题采用了类似DFS的方式，直接暴力...
*/
